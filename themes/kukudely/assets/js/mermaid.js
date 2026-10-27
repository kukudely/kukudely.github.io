const diagrams = Array.from(document.querySelectorAll('.mermaid-diagram'), (element) => ({
    element,
    source: element.querySelector('.mermaid-source').textContent,
}));

if (diagrams.length) {
    try {
        const { default: mermaid } = await import('https://cdn.jsdelivr.net/npm/mermaid@12.2.0/dist/mermaid.esm.min.mjs');
        await document.fonts.ready;
        let renderId = 0;
        const dialog = document.createElement('dialog');
        dialog.className = 'mermaid-dialog';
        dialog.setAttribute('aria-label', '图示放大预览');
        dialog.innerHTML = '<div class="mermaid-toolbar"><button class="mermaid-zoom-out" type="button" aria-label="缩小图示">−</button><button class="mermaid-zoom-reset" type="button" aria-label="恢复原始大小">100%</button><button class="mermaid-zoom-in" type="button" aria-label="放大图示">＋</button><button class="mermaid-close" type="button">关闭</button></div><div class="mermaid-dialog-canvas" tabindex="0" role="region" aria-label="完整图示，可滚动查看"></div>';
        document.body.append(dialog);
        let expanded = null;
        let rendering = false;
        let zoom = 1;
        const setZoom = (value) => {
            if (!expanded) return;
            zoom = Math.max(0.5, Math.min(2, value));
            expanded.svg.style.width = `${expanded.svg.viewBox.baseVal.width * zoom}px`;
            dialog.querySelector('.mermaid-zoom-reset').textContent = `${Math.round(zoom * 100)}%`;
        };
        dialog.querySelector('.mermaid-zoom-out').addEventListener('click', () => setZoom(zoom - 0.25));
        dialog.querySelector('.mermaid-zoom-in').addEventListener('click', () => setZoom(zoom + 0.25));
        dialog.querySelector('.mermaid-zoom-reset').addEventListener('click', () => setZoom(1));
        dialog.querySelector('.mermaid-close').addEventListener('click', () => dialog.close());
        dialog.addEventListener('close', () => {
            if (expanded) {
                expanded.svg.style.removeProperty('width');
                expanded.canvas.replaceChildren(expanded.svg);
                expanded = null;
            }
        });
        for (const { element } of diagrams) {
            element.querySelector('.mermaid-expand').addEventListener('click', () => {
                const canvas = element.querySelector('.mermaid-canvas');
                const svg = canvas.querySelector('svg');
                if (!svg || dialog.open || rendering) return;
                // 移动原节点，避免复制 SVG 导致标记和无障碍 ID 重复。
                expanded = { canvas, svg };
                dialog.querySelector('.mermaid-dialog-canvas').replaceChildren(svg);
                dialog.showModal();
                setZoom(1);
            });
        }

        const renderDiagrams = async () => {
            // 主题更新前归还节点，避免弹窗中的旧图覆盖新渲染结果。
            if (expanded) {
                expanded.svg.style.removeProperty('width');
                expanded.canvas.replaceChildren(expanded.svg);
                expanded = null;
                dialog.close();
            }
            const dark = document.documentElement.dataset.theme === 'dark';
            const theme = dark ? 'dark' : 'default';
            mermaid.initialize({
                startOnLoad: false,
                securityLevel: 'strict',
                htmlLabels: false,
                theme: 'base',
                look: 'classic',
                fontFamily: getComputedStyle(document.documentElement).getPropertyValue('--sans').trim(),
                themeVariables: {
                    darkMode: dark,
                    fontSize: '16px',
                    primaryColor: dark ? '#243449' : '#e8f0fa',
                    primaryTextColor: dark ? '#dbeafe' : '#1e40af',
                    primaryBorderColor: dark ? '#5d7590' : '#b8c8dc',
                    lineColor: dark ? '#a9b9ca' : '#64748b',
                    secondaryColor: dark ? '#1e293b' : '#f1f5f9',
                    tertiaryColor: dark ? '#182432' : '#f8fafc',
                    clusterBkg: dark ? '#182432' : '#f8fafc',
                    clusterBorder: dark ? '#526478' : '#cbd5e1',
                    edgeLabelBackground: dark ? '#101418' : '#ffffff',
                    actorBkg: dark ? '#243449' : '#e8f0fa',
                    actorBorder: dark ? '#5d7590' : '#b8c8dc',
                    actorTextColor: dark ? '#dbeafe' : '#1e40af',
                    signalColor: dark ? '#cbd5e1' : '#334155',
                    signalTextColor: dark ? '#dbeafe' : '#334155',
                    noteBkgColor: dark ? '#243449' : '#f1f5f9',
                    noteTextColor: dark ? '#dbeafe' : '#334155',
                    noteBorderColor: dark ? '#5d7590' : '#cbd5e1',
                },
                flowchart: {
                    useMaxWidth: true,
                    wrappingWidth: 240,
                    nodeSpacing: 24,
                    rankSpacing: 26,
                    padding: 12,
                },
                sequence: {
                    useMaxWidth: true,
                    wrap: true,
                    mirrorActors: false,
                    width: 110,
                    height: 40,
                    actorMargin: 32,
                    messageMargin: 28,
                    noteMargin: 8,
                    boxMargin: 8,
                },
            });

            rendering = true;
            for (const { element } of diagrams) element.querySelector('.mermaid-expand').disabled = true;
            for (const { element, source } of diagrams) {
                const stage = document.createElement('div');
                stage.style.cssText = 'position:absolute;left:-10000px;top:0;width:1200px;visibility:hidden;';
                document.body.append(stage);
                try {
                    const { svg } = await mermaid.render(`mermaid-${++renderId}`, source, stage);
                    const canvas = element.querySelector('.mermaid-canvas');
                    canvas.innerHTML = svg;
                    const diagram = canvas.querySelector('svg');
                    // 等待挂载后的布局完成，避免采用临时渲染节点的过期边界。
                    await new Promise(requestAnimationFrame);
                    if (diagram.querySelector('.nodes')) {
                        const { x, y, width, height } = diagram.getBBox();
                        const padding = 12;
                        if ([x, y, width, height].every(Number.isFinite) && width > 0 && height > 0) {
                            diagram.setAttribute('viewBox', `${x - padding} ${y - padding} ${width + padding * 2} ${height + padding * 2}`);
                        }
                    }
                    diagram.removeAttribute('height');
                    diagram.removeAttribute('width');
                    diagram.style.removeProperty('width');
                    diagram.style.removeProperty('max-width');
                    diagram.style.setProperty('--diagram-width', `${Math.ceil(diagram.viewBox.baseVal.width)}px`);
                    element.querySelector('.mermaid-expand').hidden = false;
                    element.dataset.renderedTheme = theme;
                } catch (error) {
                    console.error('Mermaid diagram could not be rendered:', error);
                } finally {
                    stage.remove();
                }
            }
            rendering = false;
            for (const { element } of diagrams) element.querySelector('.mermaid-expand').disabled = false;
        };

        // 串行渲染，避免快速切换主题时旧结果覆盖新结果。
        let pending = Promise.resolve();
        const scheduleRender = () => {
            pending = pending.then(renderDiagrams).catch((error) => {
                console.error('Mermaid rendering failed:', error);
            });
        };
        new MutationObserver(scheduleRender).observe(document.documentElement, {
            attributes: true,
            attributeFilter: ['data-theme'],
        });
        scheduleRender();
    } catch (error) {
        // 加载失败时保留代码块，读者仍可查看图示源码。
        for (const { element } of diagrams) {
            const notice = document.createElement('p');
            notice.className = 'mermaid-notice';
            notice.textContent = '图示暂未加载，以下为 Mermaid 源码。';
            element.prepend(notice);
        }
        console.error('Mermaid could not be loaded:', error);
    }
}

(() => {
  const navigation = document.querySelector('.post-navigation');
  if (!navigation) return;
  const controls = navigation.querySelector('.post-nav-controls');
  const buttons = Array.from(controls.querySelectorAll('button'));
  const header = document.querySelector('.site-header');
  const mobile = matchMedia('(max-width: 960px)');
  const wide = matchMedia('(min-width: 1360px)');
  let selected = navigation.dataset.active;
  let opened = false;

  function render() {
    navigation.dataset.active = selected;
    navigation.classList.toggle('is-open', mobile.matches && opened);
    const tabs = !mobile.matches && !wide.matches;
    if (tabs) controls.setAttribute('role', 'tablist');
    else controls.removeAttribute('role');
    for (const button of buttons) {
      const active = button.dataset.panel === selected;
      const panel = document.getElementById(button.getAttribute('aria-controls'));
      button.tabIndex = tabs && !active ? -1 : 0;
      if (tabs) {
        button.setAttribute('role', 'tab');
        button.setAttribute('aria-selected', String(active));
        button.removeAttribute('aria-expanded');
        panel.setAttribute('role', 'tabpanel');
        panel.setAttribute('aria-labelledby', button.id);
        panel.tabIndex = 0;
      } else {
        button.removeAttribute('role');
        button.removeAttribute('aria-selected');
        button.setAttribute('aria-expanded', String(mobile.matches && opened && active));
        panel.removeAttribute('role');
        panel.setAttribute('aria-labelledby', panel.querySelector('h2').id);
        panel.removeAttribute('tabindex');
      }
    }
  }

  function close(restoreFocus = false) {
    opened = false;
    render();
    if (restoreFocus) buttons.find(button => button.dataset.panel === selected).focus();
  }

  for (const button of buttons) {
    button.addEventListener('click', () => {
      opened = mobile.matches && (selected !== button.dataset.panel || !opened);
      selected = button.dataset.panel;
      render();
    });
    button.addEventListener('keydown', event => {
      if (mobile.matches || wide.matches) return;
      const index = buttons.indexOf(button);
      let next;
      if (event.key === 'ArrowRight') next = (index + 1) % buttons.length;
      if (event.key === 'ArrowLeft') next = (index + buttons.length - 1) % buttons.length;
      if (event.key === 'Home') next = 0;
      if (event.key === 'End') next = buttons.length - 1;
      if (next === undefined) return;
      event.preventDefault();
      selected = buttons[next].dataset.panel;
      render();
      buttons[next].focus();
    });
  }
  navigation.addEventListener('click', event => {
    const link = event.target.closest('.post-toc-nav a');
    if (!link || !mobile.matches) return;
    close();
    const heading = document.getElementById(decodeURIComponent(link.hash.slice(1)));
    if (heading) {
      heading.setAttribute('tabindex', '-1');
      heading.focus({ preventScroll: true });
    }
  });
  document.addEventListener('keydown', event => {
    if (event.key === 'Escape' && opened) close(true);
  });
  document.addEventListener('click', event => {
    if (opened && !navigation.contains(event.target)) close();
  });
  function resize() {
    opened = false;
    render();
  }
  mobile.addEventListener('change', resize);
  wide.addEventListener('change', resize);
  if (header) {
    const measureHeader = () => document.documentElement.style.setProperty(
      '--post-header-height', `${header.getBoundingClientRect().height}px`);
    new ResizeObserver(measureHeader).observe(header);
    measureHeader();
  }
  render();
  navigation.classList.add('is-enhanced');
})();

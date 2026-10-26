(() => {
  const links = Array.from(document.querySelectorAll('.post-toc-nav a'));
  const headings = Array.from(document.querySelectorAll('.post .content h2[id], .post .content h3[id]'));
  const header = document.querySelector('.site-header');
  if (!links.length || !headings.length) return;

  const targets = links.map(link => ({ link, id: decodeURIComponent(link.hash.slice(1)) }));
  let pending = false;
  let currentId;

  function update() {
    pending = false;
    const offset = (header ? header.getBoundingClientRect().height : 0) + 40;
    let current = headings[0];
    for (const heading of headings) {
      if (heading.getBoundingClientRect().top > offset) break;
      current = heading;
    }
    if (current.id === currentId) return;
    currentId = current.id;
    for (const { link, id } of targets) {
      if (id === currentId) link.setAttribute('aria-current', 'location');
      else link.removeAttribute('aria-current');
    }
  }

  function scheduleUpdate() {
    if (pending) return;
    pending = true;
    requestAnimationFrame(update);
  }

  document.querySelector('.post-toc-mobile').addEventListener('click', event => {
    if (event.target.closest('a')) event.currentTarget.open = false;
  });
  window.addEventListener('scroll', scheduleUpdate, { passive: true });
  window.addEventListener('resize', scheduleUpdate);
  window.addEventListener('hashchange', scheduleUpdate);
  window.addEventListener('load', scheduleUpdate);
  update();
})();

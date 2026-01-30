(() => {
  // Respect accessibility: let people opt out
  const reduceMotion = window.matchMedia?.(
    "(prefers-reduced-motion: reduce)",
  )?.matches;
  if (reduceMotion) return;

  const MAX = 6; // number of particles in circulation
  const LIFE = 666; // ms until a particle disappears
  const SPREAD = 12; // random offset around the pointer

  const container = document.createElement("div");
  container.setAttribute("aria-hidden", "true");
  container.style.cssText = `
    position: fixed;
    inset: 0;
    pointer-events: none;
    z-index: 2147483647;
    overflow: visible;
  `;
  document.documentElement.appendChild(container);

  const chars = ["😂"];
  const nodes = [];

  for (let i = 0; i < MAX; i++) {
    const el = document.createElement("div");
    el.textContent = chars[i % chars.length];
    el.style.cssText = `
      position: absolute;
      left: 0; top: 0;
      font-family: "Comic Sans MS", "Verdana", system-ui, sans-serif;
      font-size: 14px;
      opacity: 0;
      transform: translate(-9999px, -9999px) scale(1);
      will-change: transform, opacity;
      user-select: none;
    `;
    container.appendChild(el);
    nodes.push({ el, t: 0, x: 0, y: 0 });
  }

  let idx = 0;

  function spawn(x, y) {
    const p = nodes[idx];
    idx = (idx + 1) % nodes.length;

    const ox = (Math.random() * 2 - 1) * SPREAD;
    const oy = (Math.random() * 2 - 1) * SPREAD;

    p.x = x + ox;
    p.y = y + oy;
    p.t = performance.now();

    p.el.style.opacity = "1";
    p.el.style.transform = `translate(${p.x}px, ${p.y}px) scale(1)`;
  }

  function tick(now) {
    for (const p of nodes) {
      if (!p.t) continue;
      const age = now - p.t;
      if (age >= LIFE) {
        p.t = 0;
        p.el.style.opacity = "0";
        continue;
      }
      // fade out + float up, classic
      const k = age / LIFE;
      const y = p.y - k * 22;
      const s = 1 - k * 0.4;
      p.el.style.opacity = String(1 - k);
      p.el.style.transform = `translate(${p.x}px, ${y}px) scale(${s})`;
    }
    requestAnimationFrame(tick);
  }
  requestAnimationFrame(tick);

  let last = 0;
  const throttle = 16; // ~60fps
  window.addEventListener(
    "pointermove",
    (e) => {
      const now = performance.now();
      if (now - last < throttle) return;
      last = now;
      spawn(e.clientX, e.clientY);
    },
    { passive: true },
  );
})();

(function () {
  const root = document.getElementById("randomPicker");
  if (!root) return;

  const btn = document.getElementById("randomBtn");
  const out = document.getElementById("randomOut");
  if (!btn || !out) return;

  let items = [];
  try {
    items = JSON.parse(root.dataset.items || "[]");
  } catch (e) {
    return;
  }
  if (!Array.isArray(items) || items.length === 0) return;

  const variant = root.dataset.variant || "generic";

  function link(url, text) {
    const a = document.createElement("a");
    a.href = url;
    a.textContent = text;
    if (url && url.startsWith("https:")) {
      a.target = "_blank";
      a.rel = "noopener";
    }
    return a;
  }

  function pick() {
    const item = items[Math.floor(Math.random() * items.length)];
    out.textContent = "";
    if (variant === "quotes") {
      out.append('"' + item.quote + '" — ', link(item.url, item.author));
    } else {
      out.append(link(item.url, item.title));
    }
  }

  btn.addEventListener("click", pick);
})();

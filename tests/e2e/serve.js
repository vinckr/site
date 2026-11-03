// Minimal static server for the e2e tests. Serves the built _site folder the
// way Netlify does: folders resolve to index.html, unknown paths get 404.html.
const fs = require("fs");
const http = require("http");
const path = require("path");

const root = path.resolve(__dirname, "../../_site");
const port = Number(process.env.PORT) || 4173;

const types = {
  ".css": "text/css",
  ".gif": "image/gif",
  ".html": "text/html; charset=utf-8",
  ".ico": "image/x-icon",
  ".jpg": "image/jpeg",
  ".js": "text/javascript",
  ".json": "application/json",
  ".png": "image/png",
  ".svg": "image/svg+xml",
  ".ttf": "font/ttf",
  ".txt": "text/plain",
  ".webmanifest": "application/manifest+json",
  ".webp": "image/webp",
  ".xml": "application/xml",
};

function resolve(pathname) {
  let file;
  try {
    file = path.join(root, decodeURIComponent(pathname));
  } catch {
    return null;
  }
  if (!file.startsWith(root)) return null;
  for (const candidate of [file, path.join(file, "index.html")]) {
    if (fs.existsSync(candidate) && fs.statSync(candidate).isFile()) return candidate;
  }
  return null;
}

http
  .createServer((req, res) => {
    const { pathname } = new URL(req.url, "http://localhost");
    const file = resolve(pathname) || path.join(root, "404.html");
    if (!fs.existsSync(file)) {
      res.writeHead(404).end("not found (did the build run?)");
      return;
    }
    const status = file.endsWith("404.html") && pathname !== "/404.html" ? 404 : 200;
    res.writeHead(status, { "Content-Type": types[path.extname(file)] || "application/octet-stream" });
    fs.createReadStream(file).pipe(res);
  })
  .listen(port, "127.0.0.1", () => console.log(`Serving ${root} on http://localhost:${port}`));

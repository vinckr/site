import fs from "fs-extra";
import { markdown } from "./utils/markdown.js";

(async function () {
  console.log("Converting Markdown to HTML...");

  // markdown source
  const content = await fs.readFile("./content/links.md", "utf8");

  // converted to HTML
  const rendered = await markdown.render(content);


  const htmlFile = `<!DOCTYPE html>
  <html lang="en">
  <head>
  <meta charset="UTF-8" />
  <title>Link Collection - vinckr.com</title>
  <link rel="stylesheet" href="./public/styles/default.css">
  </head>
  <body>
  ${rendered}
  </body>
  </html>`;

//  await fs.mkdirs("./public");

  await fs.writeFile("./links.html", htmlFile, "utf8");

  console.log("HTML generated.");
})();
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
	<title>Link Collection - vinckr.com</title>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" href="./public/styles/default.css">
</head>
<body>
  <div class="nav">
    <ul>
      <li><a href="./index.html" >Home</a></li>
      <li><a href="./links.html">Links</a></li>
      <li><a href="mailto:hi@vinckr.com">Contact</a></li>
    </ul>
  </div>
  <div class="main">
    ${rendered}
  </div>
</body>
</html>`;

//  await fs.mkdirs("./public");

  await fs.writeFile("./links.html", htmlFile, "utf8");

  console.log("HTML generated.");
})();


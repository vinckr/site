module.exports = function (eleventyConfig) {
  // Optional during upgrade only
  // eleventyConfig.addPlugin(require("@11ty/eleventy-upgrade-help"));

  // Passthroughs
  eleventyConfig.addPassthroughCopy("styles");
  eleventyConfig.addPassthroughCopy("img");
  eleventyConfig.addPassthroughCopy({ public: "." });

  // Markdown setup
  const markdownIt = require("markdown-it");
  const mila = require("markdown-it-link-attributes");
  const markdownItAnchor = require("markdown-it-anchor");
  const markdownItTaskLists = require("markdown-it-task-lists");

  const md = markdownIt({ html: true });

  md.use(mila, {
    matcher(href) {
      return href.startsWith("https:");
    },
    attrs: { target: "_blank", rel: "noopener" },
  });

  md.use(markdownItAnchor, {
    permalink: markdownItAnchor.permalink.linkInsideHeader(),
    slugify: (str) => str.trim().toLowerCase().replace(/[\s]+/g, "-"),
    level: 2,
  });

  md.use(markdownItTaskLists);

  eleventyConfig.setLibrary("md", md);

  return {
    dir: {
      input: "content",
      output: "_site",
      data: "_data",
      includes: "_includes",
    },
  };
};

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

  // Add loading="lazy" and decoding="async" to all markdown images
  const defaultImageRender =
    md.renderer.rules.image ||
    function (tokens, idx, options, env, self) {
      return self.renderToken(tokens, idx, options);
    };
  md.renderer.rules.image = function (tokens, idx, options, env, self) {
    const token = tokens[idx];
    if (token.attrIndex("loading") < 0) token.attrSet("loading", "lazy");
    if (token.attrIndex("decoding") < 0) token.attrSet("decoding", "async");
    return defaultImageRender(tokens, idx, options, env, self);
  };

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

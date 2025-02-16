module.exports = function (eleventyConfig) {
  eleventyConfig.addPassthroughCopy("styles");
  eleventyConfig.addPassthroughCopy("img");
  eleventyConfig.addPassthroughCopy("js");
  eleventyConfig.addPassthroughCopy({
    public: ".",
  });

  const markdownIt = require("markdown-it");
  const mila = require("markdown-it-link-attributes");
  const markdownItAnchor = require("markdown-it-anchor");

  // Create a markdown-it instance with HTML enabled
  const md = markdownIt({ html: true });

  // Use markdown-it-link-attributes for external links
  md.use(mila, {
    matcher(href, config) {
      return href.startsWith("https:");
    },
    attrs: {
      target: "_blank",
      rel: "noopener",
    },
  });

  // Use markdown-it-anchor to generate IDs for headings
  md.use(markdownItAnchor, {
    permalink: markdownItAnchor.permalink.linkInsideHeader(), // Adds a link icon inside the header
    slugify: (str) => str.trim().toLowerCase().replace(/[\s]+/g, "-"), // Custom slugify logic
    level: 2, // Only use headings with a level of 2
  });

  // Set the markdown-it instance as the markdown library
  eleventyConfig.setLibrary("md", md);

  return {
    dir: {
      input: "content",
      output: "_site",
      includes: "_includes",
    },
  };
};

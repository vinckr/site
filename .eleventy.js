const mila = require("markdown-it-link-attributes");

module.exports = function (eleventyConfig) {
  // Passthrough copies for styles, images, and javascript
  eleventyConfig.addPassthroughCopy("styles");
  eleventyConfig.addPassthroughCopy("img");
  eleventyConfig.addPassthroughCopy("js");
  eleventyConfig.addPassthroughCopy({ public: "." });

  // Markdown configuration with mila
  const milaOptions = {
    attrs: {
      target: "_blank",
      rel: "noopener",
    },
  };

  eleventyConfig.setLibrary("md", (mdLib) => mdLib.use(mila, milaOptions));

  // Eleventy directory settings
  return {
    dir: {
      input: "content", // process markdown files in content
      output: "_site",
      includes: "_includes",
    },
  };
};

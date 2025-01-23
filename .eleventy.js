module.exports = function (eleventyConfig) {
  eleventyConfig.addPassthroughCopy("styles");
  eleventyConfig.addPassthroughCopy("img");
  eleventyConfig.addPassthroughCopy("js");
  eleventyConfig.addPassthroughCopy({
    public: ".",
  });

  var mila = require("markdown-it-link-attributes");
  var md = require("markdown-it")({ html: true });

  md.use(mila, {
    matcher(href, config) {
      return href.startsWith("https:");
    },
    attrs: {
      target: "_blank",
      rel: "noopener",
    },
  });

  eleventyConfig.setLibrary("md", md);

  return {
    dir: {
      input: "content",
      output: "_site",
      includes: "_includes",
    },
  };
};

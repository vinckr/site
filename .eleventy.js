module.exports = function (eleventyConfig) {
  eleventyConfig.addPassthroughCopy("styles");
  eleventyConfig.addPassthroughCopy("img");
  eleventyConfig.addPassthroughCopy("js");
  eleventyConfig.addPassthroughCopy({
    public: ".",
  });

  const mila = require("markdown-it-link-attributes");
  const milaOptions = {
    shouldOpenExternal: (href) => href.match(/^https?:\/\/(?!vinckr\.com)/),
    attrs: {
      target: "_blank",
      rel: "noopener",
    },
  };

  eleventyConfig.amendLibrary("md", (mdLib) => mdLib.use(mila, milaOptions));

  return {
    dir: {
      input: "content",
      output: "_site",
      includes: "_includes",
    },
  };
};

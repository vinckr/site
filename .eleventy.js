module.exports = function (eleventyConfig) {
  eleventyConfig.addPassthroughCopy("styles.css");
  eleventyConfig.setBrowserSyncConfig({
    open: true, // Browser will automatically open a new tab to preview your site when you start the Eleventy server.
  });

  return {
    dir: {
      input: "blog",
      output: "_site",
      includes: "_includes",
      data: "_data",
    },
  };
};

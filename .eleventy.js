module.exports = function (eleventyConfig) {
  eleventyConfig.addPassthroughCopy("styles"); // copy CSS styles to _site/styles folder
  eleventyConfig.addPassthroughCopy("img"); // copy images to _site/img folder
  eleventyConfig.addPassthroughCopy("js"); // copy javascript files to _site/js folder
  eleventyConfig.addPassthroughCopy({
    public: ".",
  }); // copy all files in public folder to _site folder

  return {
    dir: {
      input: "content", // only process markdown files in content folder
      output: "_site",
      includes: "_includes",
    },
  };
};

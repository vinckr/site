module.exports = function (eleventyConfig) {
  eleventyConfig.addPassthroughCopy("styles"); // copy all styles to _site folder
  eleventyConfig.addPassthroughCopy("img"); // copy all images to _site folder
  eleventyConfig.addPassthroughCopy({
    input: "public",
    output: "_site",
    callback: (file) => {
      let newPath = path.join("_site", path.basename(file.inputPath));
      if (fs.existsSync(newPath)) return false; // Skip copying if the file already exists in the output folder

      return true;
    },
  }); // copy all files in public folder to _site folder
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

const { test: base, expect } = require("@playwright/test");

// A test fails if the browser logs an error for this site, like a script
// exception or a missing stylesheet. Third-party requests (the Mixcloud
// embeds) are blocked so the tests don't depend on other people's servers.
const test = base.extend({
  page: async ({ page, baseURL }, use) => {
    const origin = new URL(baseURL).origin;
    const errors = [];
    await page.route(
      (url) => url.protocol.startsWith("http") && url.origin !== origin,
      (route) => route.abort(),
    );
    page.on("pageerror", (err) => errors.push(err.message));
    page.on("console", (msg) => {
      if (msg.type() === "error" && msg.location().url.startsWith(origin)) {
        errors.push(`${msg.text()} (${msg.location().url})`);
      }
    });
    await use(page);
    expect(errors, "browser console errors").toEqual([]);
  },
});

module.exports = { test, expect };

const { defineConfig, devices } = require("@playwright/test");

const port = 4173;

// Tests run against the production build in _site, so build first:
// `npm run test:e2e` (or `make test`) does both.
module.exports = defineConfig({
  testDir: "tests/e2e",
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  reporter: process.env.CI ? [["github"], ["html", { open: "never" }]] : "list",
  use: {
    baseURL: `http://localhost:${port}`,
  },
  projects: [{ name: "chromium", use: { ...devices["Desktop Chrome"] } }],
  webServer: {
    command: "node tests/e2e/serve.js",
    url: `http://localhost:${port}`,
    env: { PORT: String(port) },
    reuseExistingServer: !process.env.CI,
  },
});

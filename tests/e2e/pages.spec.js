const fs = require("fs");
const path = require("path");
const { test, expect } = require("./fixtures");

// Every page Eleventy puts in the sitemap, as a path like /blog/movies/.
const sitemap = fs.readFileSync(path.join(__dirname, "../../_site/sitemap.xml"), "utf8");
const pages = [...sitemap.matchAll(/<loc>([^<]+)<\/loc>/g)].map((m) => new URL(m[1]).pathname);

test("sitemap lists the main pages", () => {
  expect(pages).toEqual(expect.arrayContaining(["/", "/blog/", "/music/", "/feed/", "/links/"]));
});

for (const pathname of pages) {
  test(`${pathname} renders and its links resolve`, async ({ page, request }) => {
    const res = await page.goto(pathname);
    expect(res?.status()).toBe(200);
    await expect(page.locator("main > header h1")).not.toBeEmpty();

    // Structured data is only useful to search engines if it is valid JSON.
    for (const json of await page.locator('script[type="application/ld+json"]').allTextContents()) {
      expect(() => JSON.parse(json), json.slice(0, 300)).not.toThrow();
    }

    // Every same-origin link, image, script and stylesheet must exist.
    const urls = await page
      .locator("a[href], img[src], link[href], script[src]")
      .evaluateAll((els) => els.map((el) => el.href || el.src));
    const origin = new URL(page.url()).origin;
    const internal = new Set(
      urls
        .map((u) => new URL(u))
        .filter((u) => u.origin === origin)
        .map((u) => u.pathname),
    );
    for (const target of internal) {
      const res = await request.get(target);
      expect.soft(res.status(), `${target} linked from ${pathname}`).toBe(200);
    }
  });
}

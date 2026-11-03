const fs = require("fs");
const path = require("path");
const { test, expect } = require("./fixtures");

const readData = (name) =>
  JSON.parse(fs.readFileSync(path.join(__dirname, "../../content/_data", `${name}.json`), "utf8"));

test("home page renders", async ({ page }) => {
  await page.goto("/");
  await expect(page).toHaveTitle("hello · vinckr");
  await expect(page.getByRole("img", { name: "Vincent" })).toBeVisible();
});

test("stylesheet is compiled by Tailwind, not the raw source", async ({ request }) => {
  const res = await request.get("/styles/tailwind.css");
  expect(res.status()).toBe(200);
  expect(await res.text()).not.toContain('@import "tailwindcss/');
});

test("primary navigation reaches every section", async ({ page, baseURL }) => {
  await page.goto("/");
  const nav = page.getByRole("navigation", { name: "Primary" });
  for (const [name, pathname] of [
    ["music", "/music/"],
    ["blog", "/blog/"],
    ["feed", "/feed/"],
    ["links", "/links/"],
    ["home", "/"],
  ]) {
    await nav.getByRole("link", { name, exact: true }).click();
    await expect(page).toHaveURL(baseURL + pathname);
  }
});

test("unknown paths get the custom 404 page", async ({ request }) => {
  const res = await request.get("/does-not-exist/");
  expect(res.status()).toBe(404);
  expect(await res.text()).toContain("404 - page not found");
});

// The list pages are generated from content/_data, so a broken or emptied
// data file shows up here as a missing list or a dead random button.
const lists = [
  { variant: "movie", pathname: "/blog/movies/", entry: "#list > li" },
  { variant: "book", pathname: "/blog/books/", entry: "#list > li" },
  { variant: "wikipedia", pathname: "/blog/wikipedia/", entry: "#list > li" },
  { variant: "quotes", pathname: "/blog/quotes/", entry: "details figure" },
];

for (const { variant, pathname, entry } of lists) {
  test(`${variant} list shows every entry and a random pick`, async ({ page }) => {
    const items = readData(variant);
    await page.goto(pathname);
    await expect(page.locator(entry)).toHaveCount(items.length);

    await page.getByRole("button", { name: `random ${variant} recommendation` }).click();
    const pick = page.locator("#randomOut a");
    await expect(pick).toHaveCount(1);
    expect(items.map((item) => item.url)).toContain(await pick.getAttribute("href"));
  });
}

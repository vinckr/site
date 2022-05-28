---
slug: "hello-world2"
date: "2022-03-24"
title: "Building my personal page with GatsbyJS"
description: "my second blogpost. wow."
---

gatsby-cli didnt work, gave up after 5 minutes to get it to work, npm is doing fine

`npx gatsby new vinckr https://github.com/gatsbyjs/gatsby-starter-blog`

starting with the starter blog but I decide to go back to the [basic starter template](https://gatsbystarterdefaultsource.gatsbyjs.io/).

Gatsby has a number of good [starters](https://www.gatsbyjs.com/starters) for all kinds of CMSs or types of websites, but I like to start with the very basic template and build on top of it to really grok how things work in Gatsby.

## Code reference

This is some boilerplate code for my own reference, this helps me to find things again.

#### Static Image

```js
    <StaticImage
      src="../images/hello.png"
      formats={["auto", "png"]}
      alt="I hope you had a good time."
      style={{ marginBottom: `1.45rem` }}
    />
```

#### Link

```js
      <Link to="/some-route/">Some link text</Link>
```

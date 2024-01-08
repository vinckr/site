---
layout: layouts/page.njk
tags: blog
title: How to use Vale as free and customizable open source style checker
---

Grammarly provides a great service that has changed my writing workflow. But I always wondered if there was a better - and more secure - way than firing up a browser and copy-pasting my text. Also I had no options to add custom rules, and seeing the mistakes is sometimes hard when your text contains false positives due to code examples and markdown syntax.

## How I use Vale as style linter for my personal website

[Vale](https://github.com/errata-ai/vale) is a syntax-aware linter for prose and my favourite open source alternative to Grammarly, used by companies like Gitlab, Linode, brew and more. If you want to support its development, you can [donate here](https://opencollective.com/vale)! You can install it locally and run it in your IDE, browser, text editor or via the CLI. I prefer running it through the CLI, as I just want to do a check once or twice before submitting my text or when making bigger changes. A continuous spell and style check through a CI like GitHub Actions is also possible.

Vale contains two main components:

- the `vale.ini`: configuration file.
- Styles: styles for your text.
  - Vocab: contains words you want to excempt from the spellcheck.

Check [the documentation](https://docs.errata.ai) for more detailed information on all features. In the following I will give a short step-by-step guide that will be enough for basic usage.

[Installation on MacOs](https://vale.sh/docs/) (when using brew): `brew install vale`. Check if Vale has been installed: `vale -h`.

Create a `.vale.ini` configuration file in your root folder. You can also add it to a folder or project:

```ini
StylesPath = vale/styles
# This tells Vale where to look for styles

Vocab = Blog
# This tells Vale which Vocab dictionary to use, i only have one, but you can have multiple

[*.md]
BasedOnStyles = vinckr
# This tells Vale to check all markdown files
# based on the styles in StylesPath/vinckr

Vale.Spelling = NO
# I don't want Vale to check spelling
```

Create a folder for your styles `mkdir -p vale/styles`.  
Put the styles in this folder. I started with the [boilerplate styles](https://github.com/errata-ai/vale-boilerplate/tree/master/styles) and removed a lot of them. Use whatever styles suit you, I recommend to start with a few and add more over time as you develop your unique style.

You can now add custom vocabulary if needed. Inside `styles` you can find the Vocab folder containing `acccept.txt` and `reject.txt`; if not create them. I don't use Vale to check spelling at the moment.

Now you can run Vale with `vale yourdocument.md` for a single document or for example to lint all markdown files in the "content" folder: `vale content/**/*.md`

You can also embed Vale in your CI like GitHub actions, IDE, text editor, or other tools. See [the Vale documentation](https://vale.sh/docs/) for more information.

---

> Let me know if you have any questions about this guide or if something is unclear or wrong!

---

If you know a way to exclude markdown links from the Vale check let me know. My problem is that it's checking all the links texts, which I often dont control or want to edit. Right now my "solution" is to remove them and then do the style check. Let me know if you know a better way!

- [Vale documentation](https://vale.sh/docs/)
- [Vale in Gitlab Documentation testing](https://docs.gitlab.com/ee/development/documentation/testing.html#vale)

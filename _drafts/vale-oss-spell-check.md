# How to use Vale as free, private and customizable open source alternative to Grammarly

Grammarly provides a great service that has changed my writing workflow. But I always wondered if there was a better way than firing up a browser and copy-pasting my text. Also I had no options to add custom rules, seeing the real errors is sometimes hard when your text also contains code examples and markdown syntax.

https://github.com/errata-ai/vale

Vale is a great open-source alternative to Grammarly, used by companies like Gitlab, Linode, brew and more. You can install it locally and run it in your IDE, browser, text editor or via the CLI. I prefer running it through the CLI, as I just want to do a check once or twice before submitting my text. A continous spell and style check is also possible.
Vale contains two main components - or rather three:

- the vale.ini
  Your main configuration file
- Styles
  The styles for your text - for example "Avoid future tense, use present instead" or "Avoid sentences longer than X words"
- Vocab
  This contains words you want to excempt from the Spellcheck, for example product names or people names that come up a lot.

https://docs.errata.ai/

Check the documentation for more detailed information on all features. In the following I will give a short step-by-step guide that will be enough for basic usage.

https://docs.gitlab.com/ee/development/documentation/testing.html#vale

Installation on MacOs is straigthforward:
https://docs.errata.ai/vale/install

`brew install vale`

Check if Vale installed correctly.

vale -h

add a
.vale.ini
to your root, in my case (MacOS) this is my user folder,
you can alternatively also add it to a folder/project individually if those require custom or different styles.

```ini
StylesPath = vale/styles

Vocab = Blog

[*.md]
BasedOnStyles = Vale, write-good
```

Create a folder for your styles
`mkdir -p vale/styles`

In this folder we are going to put our styles, I started with the boilerplate styles:
https://github.com/errata-ai/vale-boilerplate/tree/master/styles
Then I added the gitlab docs styles
https://gitlab.com/gitlab-org/gitlab/-/tree/master/doc/.vale/gitlab

```diff
- BasedOnStyles = Vale, write-good
+ BasedOnStyles = Vale, write-good, gitlab
```

We are going to change the Vocab to our custom vocabulary `Ory`

```diff
- Vocab = Blog
+ Vocab = Ory
```

Inside the Vocab should be a `acccept.txt` and `reject.txt`; if not create them.

add the following to `acccept.txt`:

```
Ory
Kratos
Keto
Hydra
Oathkeeper
```

Now these unique names will be ignored by the spellcheck. One item per line.

Finally add an alias for .mdx files, so they get treated like .md files.

```diff
+ [formats]
+ mdx = md
```

Now you can run Vale like so:

`vale yourdocument.md`
or
`vale vale src/markdown/blog/some-blogpost.mdx`

You can also embedd Vale in your IDE, text editor or other tools, see the documentation for more information
https://docs.errata.ai/

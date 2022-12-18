# 🍑 - "A Static Site" Builder

This is my personal website, there are many like it but this one is mine.

It is built using [html/template](https://pkg.go.dev/html/template) and [gomarkdown](https://github.com/gomarkdown/markdown).

## Howto

Run to show all available commands:

```bash
make help

u                format, encrypt drafts, build HTML pages and commit to git
format           format the source code
test             run all linters
build            build HTML without committing
links-sort       sort links and delete duplicates
links            generate markdown syntax links from urls.txt
encrypt-drafts   encrypt files in _drafts folder
decrypt-drafts   decrypt files in drafts folder
```

## TODO Launch

- [ ] dynamically build navigation
- [ ] mobile/desktop view
- [ ] proper struct for global and site config
  - [ ] art (+pic-host)
  - [ ] music (+ mixcloud embed + 5 tracks blog)
- [ ] new sitemap/navigation
- [ ] page category?

Link-gen currently breaks on empty lines, 404 links, direct links to .png, .pdf ect.

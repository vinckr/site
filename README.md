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

### links

```bash
# sort links in url-unsorted.txt
make links-sort
# generate markdown link list with titles
make links
```

## TODO Launch

- [ ] dynamically build navigation
- [ ] mobile/desktop view
- [ ] proper struct for global and site config
  - [ ] art (+pic-host)
  - [ ] music (+ mixcloud embed + 5 tracks blog)
- [ ] new sitemap/navigation
- [ ] page category?

## Links

Currently breaks on:

- empty lines
- 404 links
- direkt links to .png, .pdf ect.

## Structure:

welcome page -> hand-written HTML

blog -> generated from markdown using mdtohtml

content-pages -> generated from markdown using mdtohtml

galleries?

[![](https://mermaid.ink/img/pako:eNpVzz0LwjAQBuC_Es4lBS24ZlCs9QtURAWHxiG0pw1tkpKmioj_3Vg62JvunnuX9w2pyRAY3K2ocnKOuSZ-ZskFy9QovJLRaEIiGpXmHnS_luZ0K3VR92xJZ9YFv_tf1_SID4nPfnZFd00t02mHUYtxcjC1I4Px9V8XSRiGPdl0uX2rMASFVgmZ-R7vn3BwOSrkwPyaCVtw4Prjc6Jx5vTSKTBnGxxCU2XCYSyFr6-A3URZ4-cLtzBNKw)](https://mermaid.live/edit#pako:eNpVzz0LwjAQBuC_Es4lBS24ZlCs9QtURAWHxiG0pw1tkpKmioj_3Vg62JvunnuX9w2pyRAY3K2ocnKOuSZ-ZskFy9QovJLRaEIiGpXmHnS_luZ0K3VR92xJZ9YFv_tf1_SID4nPfnZFd00t02mHUYtxcjC1I4Px9V8XSRiGPdl0uX2rMASFVgmZ-R7vn3BwOSrkwPyaCVtw4Prjc6Jx5vTSKTBnGxxCU2XCYSyFr6-A3URZ4-cLtzBNKw)

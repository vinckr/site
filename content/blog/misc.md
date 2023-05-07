---
layout: page.njk
tags: blog
title: Mixed cheatsheet
---

A mixed collection of things I had to look up at some point and collect here for future reference.

- [Mermaid Cheat Sheet](https://jojozhuang.github.io/tutorial/mermaid-cheat-sheet/)

## Mixed/UNIX

- `history | awk '{print $2}' | sort | uniq -c | sort -rn | head`: see which commands you run often.
- `curl ipinfo.io`: external ip address using ipinfo.io
- `lsof -iTCP -sTCP:LISTEN -n -P`: show all active ports.
- `kill -9 <PID>`: kill process for a port.
- `dig WWW.VINCKR.COM +nostats +nocomments +nocmd`: DNS lookup.
- `ps uax`: networks.
- `chmod +x name_of_file.sh`: make bash file executable.

## CSS

- `div > p`: child selector selects all elements that are the children of a specified element ([See also](https://www.w3schools.com/css/css_combinators.asp))

## Docker

- `docker kill $(docker ps -q)`: stop all Docker containers.
- `docker rm $(docker ps -a -q)`: remove all Docker containers.
- `docker rmi $(docker images -q)`: remove all Docker images.

## git town

Aliases for git-town, add to git config:

```
git config --global alias.append "town append"
git config --global alias.hack "town hack"
git config --global alias.kill "town kill"
git config --global alias.repo "town repo"
git config --global alias.ship "town ship"
git config --global alias.sync "town sync"
git config --global alias.new-pr "town new-pull-request"
git config --global alias.prepend "town prepend"
git config --global alias.prune-branches "town prune-branches"
git config --global alias.rename-branch "town rename-branch"
```

## regex

Replace newlines:

```
\r\n?|\n
```

## markdown

Documenting some lesser know markdown features of my blog, curtesy of [CommonExtensions](https://pkg.go.dev/github.com/gomarkdown/markdown/parser#CommonExtensions) and [Footnotes](https://pkg.go.dev/github.com/gomarkdown/markdown/parser#Footnotes)

- Ignore emphasis markers inside w*ord*s

- Add a footnote[^1]:

```
Here is a footnote reference[^1] and some more text.

[^1]: Here is the footnote.
```

- Markdown tables

```
| Id  | Name       | Signed-In |
| --- | ---------- | --------  |
| 1   | John Smith | Yes       |
| 2   | John Doe   | No        |
```

| Id  | Name       | Signed-In |
| --- | ---------- | --------- |
| 1   | John Smith | Yes       |
| 2   | John Doe   | No        |

[^1]: Bumblebees can fly, and so can you 🐝

- Strikethrough

~~Noone should read this~~

`~~Noone should read this~~`

- Definition lists

Term 1

: Definition 1

Term 2 with _inline markup_

: Definition 2

        { some code, part of Definition 2 }

    Third paragraph of definition 2.

```
Term 1

: Definition 1

Term 2 with _inline markup_

: Definition 2

        { some code, part of Definition 2 }

    Third paragraph of definition 2.
```

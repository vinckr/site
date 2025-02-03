# vinckr.com

This is my personal website, there are many like it but this one is mine.

It is built using [11ty](https://www.11ty.dev/) and [Tailwind CSS](https://tailwindcss.com/).

Feel free to open an issue or write me if you see something that looks odd or anything else you think is uh not great!

## Development

Show all available commands:

```bash
make help
```

Run the dev server:

```bash
make dev
# or
npm run start
```

## Commit messages

Note: the git history prior to [6d2454f](https://github.com/vinckr/vinckr.com/commit/6d2454f827bbffddc43dc0a6814de5064cbdcedd) is a mess so feel free to ignore that! I'm striving to keep the history clean using [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) going forward, but since I am the only one contributing to this I am not always motivated to so.

`fix`: patches a bug

`feat`: introduces a new feature

`docs`: changes to the README and other meta docs

`chore`: fixing small errors in articles, typos, archiving drafts, and similar (uninteresting) chores

`refactor`: renames something or moves a file

not part of conventional commits and specific to this site:

`weblog`: new entry to the monthly weblog

`article`: new article

`edit`: any edits to existing artices, weblogs, lists or other text

## TODO

In order of importance?! plans change a lot here so don't count on me doing this actually.

### next

- [ ] redesign landing page
- [x] change lists/pages format to json, build dedicated `list` template
  - [x] wikipedia
  - [x] movies
  - [x] books
  - [x] quotes
  - [x] add random movie/book button!
- [ ] weblog
  - [ ] automate format imdb links and add to movie json (also for reviews)
  - [ ] automate format goodreads links and add to books json (also for reviews)
  - [ ] automate format wiki links and add to wiki json
  - [ ] move future weblogs to a json
  - [ ] move past weblogs to a json
- [x] markdown text anchor links
- [x] markdown text anchor links, check how to do it [here](https://github.com/11ty/eleventy/issues/1593)
- [ ] add cool 90s cursor effects
  - [ ] add tons of other effects ideally without impacting performance
- [ ] order everything in /links in Audio and below
  - [ ] move /links to a json as well
  - [ ] restructure /links to be less insane
- [ ] add published date
  - [ ] add edited date
- [ ] 404 page
  - [ ] other errors?

## TODO nice to have

- [ ] change goodreads to something else, maybe worldcat
- [ ] move to my own static site builder again? or dynamic?
- [ ] markdown footnotes
- [ ] Check text with vale rules
- [ ] add better/more vale rules
- [x] Check if links are alive: https://github.com/tcort/markdown-link-check, next check in March 2024.
  - [ ] automate as cron job: [github action](https://github.com/gaurav-nelson/github-action-markdown-link-check) is in in maintenance mode and I could not figure out how to deal with all the exceptions without a lot of hassle, [the followup project](https://github.com/UmbrellaDocs/linkspector) still needs a github action, check again sometime (or contribute it you lazy bum!)

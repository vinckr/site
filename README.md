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

Note: the git history prior to [6d2454f](https://github.com/vinckr/vinckr.com/commit/6d2454f827bbffddc43dc0a6814de5064cbdcedd) is a mess so feel free to ignore that (actually it's not really perfect after as well but alas). I'm striving to keep the history clean using [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) going forward, but since I am the only one contributing to this I am not always motivated to so.

`fix`: patches a bug

`feat`: introduces a new feature

`docs`: changes to the README and other meta docs

`chore`: fixing small errors in articles, typos, archiving drafts, and similar (uninteresting) chores

`refactor`: renames something or moves a file

not part of conventional commits and specific to this site:

`blog`: new article/page/content

`edit`: any edits to existing artices, weblogs, lists or other text

`feed`: updates to my feed

## Style

This website uses the [cs16.css](https://github.com/ekmas/cs16.css), more about this [here](https://cs16.samke.me/).
Thanks a lot @emkas for creating this!

#### Page-level style overrides

> author note: maybe this is not needed I am not sure, didn't try it out...

The site uses **cs16.css** globally for theming and components.  
For pages that require custom adjustments, add a scoped override file and import it after `tailwind.css`.

Example:  
`src/pages/account/account.overrides.css`

```css
@layer components {
  .page-account .cs-btn {
    padding: 6px 10px;
    font-weight: 600;
  }
}
```

Include the override in Nunjucks template:

```
{% extends "layouts/base.njk" %}
{% block head %}
  <link rel="stylesheet" href="/assets/tailwind.css">
  <link rel="stylesheet" href="/pages/account/account.overrides.css">
{% endblock %}
{% block content %}
  <main class="page-account">
    <!-- page content -->
  </main>
{% endblock %}
```

Each override file should:

- Use @layer components so Tailwind’s purge keeps it.
- Scope rules to a unique wrapper (e.g., .page-account) to avoid global bleed.
- Never redefine cs16 variables directly; override component rules instead.

## Draft pipeline

Turns the stuff I save to Signal "Note to Self" into website material.

```sh
make draft                      # pull new Signal notes and process them
make draft-file FILE=notes.txt  # same, but read from a text file instead
make draft-dry                  # show what would happen, write nothing
make draft-apply                # move the finished feed block into content/feed.md
make draft-redo                 # reprocess every stored note, ignoring the last-run marker
```

### Where material goes

| Input | Destination |
| --- | --- |
| Goodreads link | `content/_data/book.json` (written directly) |
| IMDb link | `content/_data/movie.json` (written directly) |
| Recipe link | `content/blog/recipes.md`, "Recipe links" section |
| Wikipedia article about a film or book | resolved to IMDb/Goodreads, **staged for review** |
| Pasted recipe text | staged, needs formatting by hand |
| `Title 4/5` on its own line | "Movie ratings" section, formatted for `movies.md` |
| Instagram or Bandcamp link | "Filtered out" section of the draft, kept out of the feed |
| Anything else with a link | feed block in the draft file |
| Prose with no link | "Text" section of the draft file |

Books, movies and recipe links are written straight into the tracked files, so
`git diff` is the review step. The feed block, resolved links and loose text land
in `_staging/draft-<date>.md`, which is gitignored.

Links resolved from another source are never written into the curated lists
automatically — a search match can pick the wrong edition, so they wait in the
"Needs review" section of the draft file. They are rendered there as JSON
fragments in the exact shape of the data files, grouped per destination, so they
can be pasted in without reformatting:

```json
  {
    "title": "Stalker",
    "url": "https://www.imdb.com/title/tt0079944/"
  },
```

Where each one came from is listed underneath the block rather than inside it,
so the paste stays clean. Run `make data-sort` afterwards to re-sort.

The tool never writes anything inside `_drafts`.

### Host filter

Instagram and Bandcamp links are kept out of the feed block. They are not thrown
away — they land in a "Filtered out" section of the draft, so the occasional one
worth keeping can be moved up into the feed block by hand. YouTube is
deliberately **not** filtered; those are decided case by case.

To filter another site, add it to `skipHosts` in `cmd/draft/classify.go`.
Subdomains are matched too, so `instagram.com` also covers `de.instagram.com`.

### Input format

A pasted Signal chat export is read as-is. Date separators set the day for
everything beneath them, and the `HH:MM` line under each message closes it:

```
Fri 10 Apr

http://beigeboxoss.com/
14:44

Sat 11 Apr

https://archive.org/details/@aadam_jacobs_collection
17:40
```

Recognised separators include `Fri 10 Apr`, `13 Mar 2026`, `Mar 13, 2026`,
`13/03/2026`, `Today` and `Yesterday`. When the year is missing it is inferred as
the most recent one that is not in the future, so `25 Dec` read in October means
last December.

A plain list of links with no date headers is still fine — everything is filed
under today.

### Feed workflow

Feed links are grouped under the day they were posted, newest day first:

```md
### 11-04-2026

- [Internet Archive](https://archive.org/details/@aadam_jacobs_collection) - 

### 10-04-2026

- [BeigeBox](http://beigeboxoss.com/) - 
- [The Update, The Vent, and The Disaster](https://randsinrepose.com/archives/...) - 
```

Write the comments, delete the links you don't want, then `make draft-apply`
prepends every day in the block to `content/feed.md` above the newest entry.

### Movie ratings

A note that is just a title and a score, like `Horrible Bosses 3/5`, is treated
as a review rather than loose text. The title is looked up on IMDb and the draft
offers a line ready to paste into the reviews list in `content/blog/movies.md`:

```md
- [Horrible Bosses](https://www.imdb.com/title/tt1499658/) 3/5
- [Bugonia](https://www.imdb.com/title/tt12300742/) 4/5
```

Nothing is written to `movies.md` automatically, so scores stay yours to change.
When the IMDb search lands on a different title than the one typed, or finds
nothing at all, the draft says so underneath instead of quietly using the wrong
film.

### Signal setup

`signal-cli` is the only external dependency and is not installed by default:

```sh
brew install signal-cli
signal-cli link -n "vinckr.com draft pipeline"
```

That prints a `sgnl://linkdevice?...` URI — turn it into a QR code and scan it
with Signal on your phone (Settings → Linked devices):

```sh
brew install qrencode
signal-cli link -n "vinckr.com draft pipeline" | qrencode -t utf8
```

Then set your number in `.env`:

```sh
SIGNAL_ACCOUNT=+49123456789
```

Two things to know about linked devices:

- Signal only delivers messages to this device that were sent **after** linking.
  Anything already sitting in Note to Self won't arrive — paste that history into
  a text file once and use `make draft-file`.
- Each message is delivered exactly once, so `make draft` appends everything it
  receives to `_staging/notes.jsonl` and processes from there. Don't delete that
  file unless you mean to lose unprocessed notes.

### State

`_staging/state.json` records which notes have already been drafted, by
fingerprint rather than by date. That is what "everything since I last ran this"
means, and it is why importing an old export still works: historical dates are
not mistaken for material that was already handled.

`make draft-redo` processes everything in the store again. `make draft-reset`
empties the store and the processed markers, which is what you want after
changing how input is parsed.

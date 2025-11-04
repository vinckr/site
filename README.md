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

`blog`: new article/page/content

`edit`: any edits to existing artices, weblogs, lists or other text

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

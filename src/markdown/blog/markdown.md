---
slug: "/markdown/"
date: "2022-03-14"
title: "Markdown"
description: "cheatsheet for markdown"
published: true
---

This is a quick reference and testing ground. I use markdown as source files for this blog,

- [John Gruber's original spec](http://daringfireball.net/projects/markdown/) 
- [Github-flavored Markdown info page](http://github.github.com/github-flavored-markdown/).

## Lists

```markdown
1. First ordered list item
1. Another item
    - Unordered sub-list.
1. Just use 1s, the numbers work automatically
   1. Ordered sub-lis
1. And another item.

   Text that should be aligned with the above item.
```

1. First ordered list item
1. Another item
    - Unordered sub-list.
1. Just use 1s, the numbers work automatically
   1. Ordered sub-lis
1. And another item.

   Text that should be aligned with the above item.

## Links

```markdown
[An inline-style link](https://www.google.com)

[A reference-style link][arbitrary case-insensitive reference text]

[Numbers for reference-style link definitions][1]

Use the [link text itself][An inline-style link](https://www.example.com)

[A reference-style link][arbitrary case-insensitive reference text]

[Numbers for reference-style link definitions][1]

Use the [link text itself]

Reference links:

[arbitrary case-insensitive reference text]: https://www.example.com
[1]: http://example.com
[link text itself]: http://www.example.com
```

[An inline-style link](https://www.example.com)

[A reference-style link][arbitrary case-insensitive reference text]

[Numbers for reference-style link definitions][1]

Use the [link text itself]

[arbitrary case-insensitive reference text]: https://www.example.com
[1]: http://example.com
[link text itself]: http://www.example.com

## Images

```markdown
Inline-style:
![alt text](../../images/sunrise.png)

Reference-style:
![alt text][sunrise]

[sunrise]: ../../images/sunrise.png
```

Inline-style:
![alt text](../../images/sunrise.png)

Reference-style:
![alt text][sunrise]

[sunrise]: ../../images/sunrise.png

## Code and Syntax Highlighting

```markdown
`code` has `back-ticks around` it.
```

Inline `code` has `back-ticks around` it.

<pre lang="markdown"><code>```javascript
var s = "JavaScript syntax highlighting";
alert(s);
```

```python
s = "Python syntax highlighting"
print s
```
</code></pre>

```javascript
var s = 'JavaScript syntax highlighting';
alert(s);
```

```python
s = "Python syntax highlighting"
print s
```

## Tables

Tables don't look great right now, but you can use them.
Colons can be used to align columns.

```markdown
| Tables        | Hold          | Data  |
| ------------- | -----------: | :----: |
| col 3 is      | right-aligned | 123   |
| col 2 is      | centered      | 456   |
| some other    | text          | 789   |

```

| Tables        | Hold          | Data  |
| ------------- | -----------:  | :----:|
| col 3 is      | right-aligned | 123   |
| col 2 is      | centered      | 456   |
| some other    | text          | 789   |


## Blockquotes

```markdown
> Blockquotes to emulate reply text.
> This line is part of the same quote.

Some other text 

>This is a long quoted line that will still be quoted properly when it wraps. You can *put* **emphasis** and *italics* into a blockquote as well.
```

> Blockquotes to emulate reply text.
> This line is part of the same quote.

Some other text 

>This is a long quoted line that will still be quoted properly when it wraps. You can *put* **emphasis** and *italics* into a blockquote as well.

## Horizontal Rule

```markdown
A line

---
```

A line

--

## YouTube Videos

```markdown
<a href="https://www.youtube.com/watch?feature=player_embedded&v=dQw4w9WgXcQ
" target="_blank"><img src="https://img.youtube.com/vi/8AkLfYOgIrE/0.jpg"
alt="IMAGE ALT TEXT HERE" width="240" height="180" border="10" /></a>
```

<a href="https://www.youtube.com/watch?feature=player_embedded&v=dQw4w9WgXcQ
" target="_blank"><img src="https://img.youtube.com/vi/dQw4w9WgXcQ/0.jpg"
alt="IMAGE ALT TEXT HERE" width="240" height="180" border="10" /></a>

Or, in pure Markdown, but losing the image sizing and border:

```markdown
[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/dQw4w9WgXcQ/0.jpg)](http://www.youtube.com/watch?v=dQw4w9WgXcQ)
```

[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/dQw4w9WgXcQ/0.jpg)](http://www.youtube.com/watch?v=dQw4w9WgXcQ)

## Headers

```markdown
# H1
## H2
### H3
#### H4
##### H5
###### H6
```

# H1
## H2
### H3
#### H4
##### H5
###### H6

## Emphasis

```markdown
Emphasis, aka italics, with *asterisks* or _underscores_.

Strong emphasis, aka bold, with **asterisks** or __underscores__.

Combined emphasis with **asterisks and _underscores_**.

Strikethrough uses two tildes. ~~Scratch this.~~
```

Emphasis, aka italics, with _asterisks_ or _underscores_.

Strong emphasis, aka bold, with **asterisks** or **underscores**.

Combined emphasis with **asterisks and _underscores_**.

Strikethrough uses two tildes. ~~Scratch this.~~

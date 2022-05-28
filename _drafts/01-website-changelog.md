---
title: 'Building a reusable personal page'
date: 'December 14, 2021'
excerpt: 'Building a personal page in a couple of days/weeks'
---

My very first post, this explains a bit the motivation and how I build this personal page.
It will be a landing page with a text, a blog for tech/work related postings, and possibly some more featues.

Some objectives:

- Cheap, almost free
- Easy to maintain and add content
- Uses modern technologies, but starting out with basic HTML,CSS,JS. 
- Dont overcomplicate, hacks allowed.

I first register a domain, vincentkraus.com, and start a git repo for the blog

Then i fork a public repo called nextjs markdown blog, I need to learn more Nextjs, React and so on anyway, so why not give it a try. 

I change some flavour things, move files around, change some components, etc.
Halfway through the process I noticed that I first need to do the landing page in HTML and then translate it to NextJS.

I delete everything and start from scratch again. It is not easy to decide what you want to say on a personal page. Who is going to be the main audience? Since I am using this website professionally, I decide to keep it formal, but not stiff in tone on the landing page. And show a bit more variety of content in the rest of the website. Getting hung up like hell on the content, even if it is just a few sentences.
I start a new version from scratch in very basic HTML, and will build up from there.

the first html version is published 
https://github.com/vinckr/vinckr.github.io/commit/6ba8d06a627675146fcab8c405c3d533de88f349
it features an start page and a link list. 
i still have to polish the text and add a journal and articles section.
i have some ideas for the journal and also have to write at least two or three smaller articles for the article section. i could also republish some older material but i like the idea of starting from scractch - and I dont have any really significant material from before any way.

next todos
- [ ] improve start page
- [ ] write some article 
- [ ] finish blizzard article 
  (this one i have had in my head for a long time and finally started on it)
- [ ] design and code journal/article section, 
  ideally each journal entry and article is a simple rendered .md file.

--- 
worked a bit on the startpage
added new links
added text-runner to check for dead links

- move to astro

learn what is aria-label
https://www.aditus.io/aria/aria-label/

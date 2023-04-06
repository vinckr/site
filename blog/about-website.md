---
pagetitle: "About this website"
---

# About this website

The first time I built a website was around 2002/3 when I was still a teenager. It was a crude thing, a concoction of frames and Comic Sans. It was made with Dreamweaver - a no-code website builder tool, I wonder what became of it - and the content mostly was about Magic the Gathering decks and strategies. Yes, I was that kind of nerd in the 2000s: Magic, Warhammer and World of Warcraft <.<

The premise of this website was very similar to the current one: Having a space online that is fully my own and that I can fill with anything I feel like.
In any case I lost interest in Magic and subsequently in maintaining the website after a few years and it went into disrepair. There was just not much I wanted to write about publicly back then - if at all.

I think that was before the time the internet archive existed, at least I have not been able to find a version of the magic decks website anywhere.
After I had a few odd blogs on different platforms, joined a few forums, dappled in collection weird and unsual facebook groups - but I never really had my own personal space after that.
(insert picture here)

Until sometime during the surreal lockdown years, I deleted all my social media accounts, but still had a desire to write down my thoughts or communicate through several layers of irony with memes.
That was the original motivation to take up owning my own little space online again - apart from now working for a infosec company.

The first iteration of this website was build in Gatsby, because that is what is used at my work, so I was familiar with it.
It quickly became a headache to maintain, because there was just a lot of features that I would never use and that got in my way.
I did a short tour of different frameworks that would support building a simple static website out of markdown. Next.js, Astro, and Hugo were the ones that stood out for me and would use them again for a different project, but they all were much too "feature-packed" and complex to use for my simple personal website.

That is why I decided to write a simple program of my own, that creates a basic HTML page after a template, while using markdown as a content source.
What this means is you create the head and general structure of the HTML directly, add a {{ . }} (I call it the "double-spiky cyclops") where your markdown text should be rendered and that is it.
You can do some basic styling with CSS or just go with the raw, "brutalist" look.
You can even write custom HTML inside the markdown!
Anyway after a day or so of tinkering I got my own static site builder to work, and learned some Go in the process, yay!

Currently this website is living in a git repository and being built and deployed on Netlify whenever I push to that repository. This is basic stuff nowadays, but my website and me came a long way since it was first build and deployed via ftp.

## What next?

Right now it is just this collection of basic markdown files and very text content heavy with a minimal style.
I want to experiment more with other styles and designs.
I also want to attach other, more complex projects to this domain, which will most likely be running one of the fancy new frameworks mentioned above.
But the main part or driver of this blog will remain my own for the time being - there are so many things to play around with that I coudl use to render this, but I want to keep the focus on the actual text and less on what framework or new hype tech I want to try out this week.
The static site builder I am using for this blog can be found on GitHub here. I swear I will make a tutorial on how to use it soon - but you should probably not use it. Try out Hugo, its pretty cool, or look in my link collection, there should be something useful.

---
layout: page.njk
tags: blog
title: Stable Diffusion Text to Image
---

## How to use Text to Image

TODO Simple examples demonstrating how to use Stable Diffusion WebUI to generate images

TODO move to dedicated doc txt2img

## Basic Prompting

TODO move to dedicated doc prompt engineering Stable Diffusion

Prompting is basically describing the image you want to generate, using single words or short combinations of words seperated by a comma.
prompting involves a lot of iterations, tweaking and trying out different combinations until you find something that works. It can also be influenced by many other factors. what works for me is either trying out things, tweaking, and writing down what works - you can find a list of my favourite prompts/tokens later in this guide.

:::tip! play around with how you build prompts, there is no "right" or "magic" prompt. Whatever works for you works! You can use this guide as a starting point and over time develop your preferred method to build prompts.:::

Im going to show you how this could look with a simple example, an image of a duck.
a prompt can be a simple as a single word.

TODO duck picture

This one includes some more words to describe the image

TODO duck, lake, moon, night, dark picture

You can improve this image drastically by adding the phrase "a hyper detailed realistic oil painting". This

TODO a hyper detailed realistic oil painting of a duck, lake, moon, night, dark

Change the night setting to a summer one. Notice how the sunglasses is basically ignored. The order of the tokens influcences the prompt. I am not sure how exactly that works, but the earlier a token comes the more it influences the outcome.

TODO a hyper detailed realistic oil painting of a duck, sun, beach, sunglasses

now if we want the duck to wear the sunglasses, we can try this. in this case we very literally describe the image.

TODO a hyper detailed realistic oil painting of a duck wearing sunglasses, sun, beach

With most Stable Diffusion models a combination of tokens seperated by commas works well

pirate, eyepatch, pirate ship, sea shanty, wooden leg, ...

TODO pirate, eyepatch, pirate ship, sea shanty, wooden leg

(weirdly frightening outcome by the way)

Some people - and some diffusion models too! - prefer natural language, but the above more "unnatural language" style works best for me. I feel more in control.

a pirate wearing an eyepatch on a pirate ship singing a sea shanty while dancing on this wooden leg

TODO a pirate wearing an eyepatch on a pirate ship singing a sea shanty while dancing on this wooden leg

However combining both also leads to interesting results. Ultimately it is up to you what best works for you. With some images you want to generate one approach might work better, with some another. as I will say a milion times in this guide, you have to find out what works for you, use it as a guideline but not as a hard rule on how to do things.

a hyper detailed realistic oil painting of a pirate wearing an eyepatch, singing a sea shanty, dancing on this wooden leg, pirate ship

TODO a pirate wearing an eyepatch on a pirate ship singing a sea shanty while dancing on this wooden leg

TODO refer to links and tokens

### Tokens

Tokens are what you build the prompt out of. You use them to describe the content of the image and also to influcence the style and quality. There is a maximum number of tokens that you can put in a prompt in Stable Diffusion, so after some length the model will ignore your tokens, but generally you can use a lot of them. (TODO add exact number)

There are a lot of tokens that are more focused on improving the quality of the image. I want to put special emphasis on those because they are less intuitive than the others, where you mostly describe the content of the picture or the style.
You can also use a lot of the terms from photography, so if you have experience in photography, you can make use of a lot of terms from there.
For example our earlier duck picture with some photography terms: natural lighting, hard focus, film grain, hasselblad, 35 mm

TODO image of a hyper detailed realistic oil painting of a duck on a beach, shining sun, natural lighting, hard focus, film grain, hasselblad, 35mm

You can often increase the quality by just adding things like "masterpiece, best quality, extremely delicate and beautiful"

TODO image a hyper detailed realistic oil painting of a duck on a beach, shining sun, natural lighting, hard focus, film grain, hasselblad, 35mm, masterpiece, best quality, extremely delicate and beautiful

Here is a longer list of different tokesn including moods, materials, styles and of course quality tokens.
Of course the line between all of those are blurry and tokens filed under quality also have an influence on the style and other things. So take it with a grain of salt, but try things out and see what improves the output more towards what you want to create and make your own list of quality tokens.

If give you this list of "technical quality" tokens for inspiration as many of them are quite unintuitive or very technical terms:

trending on artstation, trending on CGSociety, trending on pixiv fanbox, award winning, tate artist id, professional octane render, uhd 8k, unreal engine, vray, houdini render, quixel megascans, arnold render, raytracing, cgi, 100mm, dslr, cinema4d, studio quality,behance, criterion collection, reimagined by industrial light and magic, game art, HD, 8k post-processing, stunning, cell shading, stimulating, inspired, creative genius, epic, captivating, frostbite 3 engine, cryengine, dof

Of course here is our image from before with all of these added:

TODO Image with the above tokens added

This one came out quite nice, but usually you want a more nuanced approach and add those tokens bit by bit instead of chucking them all in, but it can be fun.

Also see Emphasis!

## Negative prompts

Negative prompts are what you dont want to see in the picture. You can think of them in the same way as "positive"/normal prompts but inverted. So when the normal prompt will lead the model towards the latent space that resembles them the most, the negative prompt will go in the opposite direction. If that sounds complicated, I dont really understand it either (find out).
For practical reasons, you can use negative prompts in three ways:

- As a quality enhancer. For example if you include "cropped, worst quality, low quality, jpeg artifacts" in your negative prompt, the model will generate images that are the opposite of cropped, so perfectly in frame. They will also be great quality without weird artifacts. Of course its a bit more tricky than that but trying that out with the image from before:

TODO a hyper detailed realistic oil painting of a duck on a beach, shining sun

now the same with "cropped, worst quality, low quality, jpeg artifacts" in the negative prompt

TOOD same with negative cropped, worst quality, low quality, jpeg artifacts

Now I usually use a "standart" negative prompt, that worked well for me in the past and tweak that as I go along. But not too many changes in most cases. However negative prompts are **powerful** just as the normal or positive prompt. You can generate images only using the negative prompts and get great results as well. How you use the negative prompt is completely up to you, for me a combination of the quality negative prompt, plus some things I dont want to see in the image works best

Here is the same from before with my full standard negative prompt

lowres, bad anatomy, bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, blurry

Another one I also really like is this one, basically a collection of "bad quality" tags from anime sites.

lowres, cropped, artifacts, gif artifacts, jpeg artifacts, aliasing, scan artifacts, bleed through, crease, vhs artifacts, bad aspect ratio, blurry, censored, reference photo, reference work, screencap redraw, duplicate, error, bad reflection, extra digits, typo, lossless-lossy, lossy-lossless, image macro, md5 mismatch, moire, third-party edit, detexted, hard translated, poorly drawn, anatomical nonsense, bad anatomy, bad feet, bad hands, bad proportions, quality, worst quality, low quality, normal quality, ranguage, resized, downscaled, image sample, upscaled, waifu2x, reversed, solid color thumbnail, unmoving pattern, watermark, third-party watermark, web address, extraction

TODO image with the above negative prompt

You can of course also use the negative prompt to eliminate things you dont want to see in a picture

TODO image with a hyper detailed realistic oil painting of a duck on a beach, shining sun, good time, in the style of Greg Rutkowski and big negative

Now with the same seed and adding "stones, rocks" to the front of the negative prompt we get a beach without any rocks.

TODO same as above with stones rokcs in negative

Using negative prompts to remove things from your image mostly works very well, sometimes not. The only way to know is to try it out.

The third use for negative prompts is more creative and that is to generate the image more through the negative than the positive prompt.

For example if we put our positive prompt in the negative this is the result, which is really weird and not all that good:

TODO image with the inverted negative prompt

Its a bit better if we put a "quality" token in the positive prompt, a token that is mostly meant to improve the quality of the output

TODO image with the mixed inverted negative prompt

TODO image with the just duck negative prompt

As you can see you can get some pretty weird results with the negative prompt.
Also the different models respond differently to negative prompts, so best bet is try play around and try out different approaches.

So in summary, you can use negative prompts in many ways and can be just as creative with them as with the normal positive one, but there are three uses that are more or less straightfoward, and I would argue the quality one is almost mandatory its so effective:

- improving quality of output by putting "bad" words in the negative prompt
- influencing image content by putting things you want to remove in negative prompt
- describing image content through the negative prompt instead of the positive - so putting "the inverted thing" of what you want to see in the negative prompt. this is not the easiest technique to use but sometimes has weird and funny results.

Here are some more negative prompts focused on improving quality that might be an inspiration for you: TODO link

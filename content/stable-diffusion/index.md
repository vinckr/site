---
layout: page.njk
tags: blog
title: Stable Diffusion Guide Draft
---

## Introduction

This guide is intended for people new to image generation or "AI art" or whatever you might call it.  
The main technology used is Stable Diffusion (SD) 1.6 running interference - which means in our case generating images - on a local laptop or desktop computer

In this guide you will learn the basics of how to run a Stable Diffusion (SD) model locally, meaning on your local computer. Many of the prompting techniques and other workflows from this guide are also applicable to other models or transformers, but Stable Diffusion 1.6 is the main model used in this guide.

TODO why SD 1.6

TODO what is stable diffusion? I dont want to bore you with the technical details, here is a list of resources if you want to learn more about diffusion models:

TODO LINKS

TODO Brief overview of Stable Diffusion and its advantages for machine learning tasks

TODO Explanation of why Stable Diffusion is a useful tool for both beginners and advanced users in the field

## Getting Started

TODO rewrite & investigate how much RAM/GPU you need at minimum
To run stable diffusion locally you need a desktop or laptop computer with at least 16 GB RAM or something, I dont really know. I use SD on MacOS, so your experience may vary on windows/linux

- without user interface: [Run Stable Diffusion on your M1 Mac's GPU](https://replicate.com/blog/run-stable-diffusion-on-m1-mac)
- with user interface: [Stable Diffusion web UI repo on GitHub](https://github.com/AUTOMATIC1111/stable-diffusion-webui)

I would very strongly suggest to use the AUTOMATIC1111 Stable Diffusion web UI, its by far the best, most fully featured, and most popular UI for Stable Diffusion. It also has a lot of helper scripts built in, plus the possibillity to install further extensions. So you can go from generating basic images to very elaborate and complex setups. It is also possible to automate or speed up the whole process of generating images using the web UI or even build custom applications for others to use Stable Diffusion via the API mode.
[web UI feature showcase](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Features)  
Just look at how many features this has, amazing that the open source community - and largely one person [AUTOMATIC1111](https://github.com/AUTOMATIC1111) has built this in a few months. I am still catching up to all the features and it feels like every week or so there is something new to play around with. there are many other good UIs for running stable diffusion, and also a few that are more freaky like an infinite outpainting canvas one, but the stable-diffusion-webui is by far the best in almost all aspects (i might include a list for this at the end of the guide)

TODO links to Step-by-step instructions for installing Stable Diffusion on a local machine

- [Install and Run on NVidia GPUs](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Install-and-Run-on-NVidia-GPUs)
- [Installation on Apple Silicon](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Installation-on-Apple-Silicon)
- [Install and Run on AMD GPUs](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Install-and-Run-on-AMD-GPUs)

./stable-diffusion-webui/run_webui_mac.sh
or
cd stable-diffusion-webui && ./run_webui_mac.sh

## Terminology

- token
- steps
- model
- latent space

## How to use Text to Image

TODO Simple examples demonstrating how to use Stable Diffusion WebUI to generate images

## Basic Prompting

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

With most stable diffusion models a combination of tokens seperated by commas works well

pirate, eyepatch, pirate ship, sea shanty, wooden leg, ...

TODO pirate, eyepatch, pirate ship, sea shanty, wooden leg

(weirdly frightening outcome by the way)

Some people - and some diffusion models too! - prefer natural language, but the above more "unnatural language" style works best for me. I feel more in control.

a pirate wearing an eyepatch on a pirate ship singing a sea shanty while dancing on this wooden leg

TODO a pirate wearing an eyepatch on a pirate ship singing a sea shanty while dancing on this wooden leg

However combining both also leads to interesting results. Ultimately it is up to you what best works for you. With some images you want to generate one approach might work better, with some another. as I will say a milion times in this guide, you have to find out what works for you, use it as a guideline but not as a hard rule on how to do things.

a hyper detailed realistic oil painting of a pirate wearing an eyepatch, singing a sea shanty, dancing on this wooden leg, pirate ship

TODO a pirate wearing an eyepatch on a pirate ship singing a sea shanty while dancing on this wooden leg

TODO links

- https://www.krea.ai/

### Tokens

Tokens are what you build the prompt out of. You use them to describe the content of the image and also to influcence the style and quality. There is a maximum number of tokens that you can put in a prompt in stable diffusion, so after some length the model will ignore your tokens, but generally you can use a lot of them. (TODO add exact number)

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

## How to use Image to Image

Instead of starting out initially with just random noise (like with text2image) you use take another image as basis. this is a wildly different mode to the one described above, but the prompting works generally the same - it just uses the starting image as basis, so that has a big influence on the outsome.

Image2Image is powerful and you can do a lot of things with it. I use it mainly for this:

- refining an image i generated through text
- using an existing image as base for a new generation

combining the two techniques of txt2img and img2img really opens up what you can do with stable diffusion and using them both is mandatory for getting good results.

For now there is no guide here, I will add one once I get around to it, promise!

## How to use Inpainting

For now there is no guide here, I will add one once I get around to it, promise!

## How to use Outpainting

For now there is no guide here, I will add one once I get around to it, promise!

## Advanced Usage of Stable Diffusion

now it gets freakishly complex really fast, so hold onto your horses

you can train new models on other images, to sort of refine the base stable diffusion model. Think of it as stable diffusion is trained on all images in the internet, also the really bad quality ones. Now people create a new checkpoint (like a 7gb model file that you download) with for example high quality images. The resulting model will have a very strong bias towards those higher quality images, resulting in generally better quality outputs. Now depending on how the model is trained this can be a stronger or weaker effect.
The real magic happens when there are dozens of high quality models trained and then combined in dozens of layers, so you have a model like

> TODO Why new Checkpoints models are a gamechanger

### Model improvements

TODO rewrite  
In my experience, embeddings need a lot of care to produce good results. They will work fine with some models and produce only garbage with others. Similar to hypernetworks but I have been getting much better results with them. Checkpoints are very powerful and probably the easiest to use. Of course I have not tried all possible combinations, but overall I have been getting the best results using good checkpoint models.

#### Embeddings

TODO rewrite  
Embedding:
An embedding is the result of textual inversion. Textual inversion tries to find a specific prompt for a given model that creates images similar to your training data. You can only generate stuff the model is already capable of. An embedding is just a shorthand for a prompt which will be treated like a very precise prompt by the model. An embedding has a small filesize of ~150mb.

##### Embeddings Examples

dan-mumford, Marc_Allante, midjourney-style, moebius, arcane-style, john-blanche, guttestreker, doose-realistic, lugal-ki-en, huang-guang-jian, sakimi, wlop-style, malika-favre, dr-strange,joe-whiteford-artstyle

#### Checkpoints

TODO rewrite  
Checkpoint model (trained via Dreambooth or similar):
A checkpoint model can be used like the base Stable Diffusion 1.4 model. The training data is used to change weights in the model, so it can render images similar to the training data. Training a new model is not trivial, extra care needs to be taken to not override existing data, or it will only render images that look like the training data. A checkpoint model has a big filesize from ~4-7gb.

##### Checkpoints Examples

TODO

trigger words
most of the time its the file name!

##### VAE

Now I never really undestood what VAE actually is[1], but sometimes you have a small .vae.pt or vae....ckpt together with your checkpoint model. you need to select that VAE in the webUI or leave the selection on "Auto" which is honestly the best option in almost all cases. Many checkpoints have the VAE "baked in", so you dont have an extra file and you can just load the checkpoint and it will work perfectly.

If you notice your images are kind of dull, like drained of color and have weirdly coloured small circles on them, that is an indication that you might be using the wrong VAE.
Most of the time you wont need to worry about this, but I wanted to include it because it stumped me a lot in the beginning - and i actually spent a few weeks generating images with the wrong VAE, wondering why the color was looking so bad until I read more about the topic.

[1]In the context of stable diffusion models, VAE stands for Variational Autoencoder. Variational Autoencoders are a type of generative model that combines elements of both autoencoders and variational inference. They are commonly used for unsupervised learning tasks, including data generation and representation learning.

#### Hypernetworks

TODO rewrite  
Hypernetwork:
A hypernetwork is another layer that will be processed after the image has been rendered through the model. The hypernetwork will skew all results from the model towards your training data. A hypernetwork "changes" or rather transforms the models results. This has the advantage that every image that contains something related to your training data, will look like your training data. If you trained a specific cat, you will have a very hard time trying to get any other cat using the hypernetwork. Hypernetworks seemt to rely on keywords already known to the model. A hypernetwork has a very small filesize of ~80mb.

##### Hypernetworks Examples

TODO

## Advanced technique

### Weights

- weights influence how much the model steers the output towards that token
- () around a token increases

### mixing prompts

[duck|pig]

### X/Y inputs

### wildcards

### ControlNet

- https://github.com/AUTOMATIC1111/stable-diffusion-webui/discussions/7784

## We'll fix it in post

### Random tips

TODO rewrite

- generate a picture with the prompt that you want, be very precise. I personally generate pictures that are 512x512 initially.
- once you get a decent picture to come out of the generation, it will be used as the base "sketch" to feed to img2img.
- if you want to, increase the resolution but if you do so set the denoising to about .60
  once you have the resolution you want and everything, keep reprocessing the image with a denoising of about 0.2 - 0.3.
- if something on the image bothers you, work on it in an image editor, for example using a brush of the same colour of the what's adjacent to the detail you want to remove, or if you want to add something (like refining fingers), make sure to use the pencil with a contrasted colour (I generally use black).
  after editing, always reprocess the image with a denoising of about 0.3
  once the result satisfies you enough, use the "R-ESRGAN 4x+ Anime6B" upscaler if you want the image upscaled.

TODO rewrite

- General workflow: Generate a ton of images, look for cool poses, send to inpainting. Work on hands and bad anatomy with mask blur 4, inpaint at full resolution, masked content: original, 32 padding, denoise 0.2-0.4 for small changes, 0.55-0.75 for large changes. Once you have anatomy and hands nailed down, move on to cosmetic changes to clothing, then faces. I cranked up the number of steps for faces, no idea if that actually helps.

### Masking

TODO

### Upscaling

#### Real-ESRGAN

https://github.com/xinntao/Real-ESRGAN
python inference_realesrgan.py -n RealESRGAN_x4plus -i inputs --face_enhance --fp32
python inference_realesrgan.py -n RealESRGAN_x4plus -i inputs --fp32

## Training your own

### Optimising Dreambooth training for 100% successful renders

TODO rewrite & test

- At least 20 close portrait shots, At least 4 middle distance torso shots, At least 4 long distance full length body shots, At least 1 seated shot.
- Stick to well focussed natively high res images.
- For more variety in your Dreambooth models, try to use varied environments, lighting, hair styles, expressions, poses, angles and distance from subject in your images.
- To avoid random hands stuck to the head in your renders, remove all images with hands too near or touching the head.
- To avoid the fisheye lens effect renders, remove all selfie images.
- To avoid white borders in your renders, make sure your images contain no viewable cropped borders.
- To avoid unnatural excessive blur, make sure the images contain no fake heavy depth of field or bokeh.
- I edited the face in training images because it was causing ugly artifacts or completely breaking the model in a lot of cases
- Class name should be "a **\_**" We can't do this while auto generating in google colab. So I'd suggest you generate the class images yourself because we can add an "a" to your chosen class word increasing the probability of rendering what you want and not a group, book cover, comic page and all the other random stuff that using the class word alone can produce. I used 60 training images, 2000 "person" class images and trained for 6000 steps. With a well trained model, it takes SD less effort to get a good likeness, freeing resources that can be spent on a much higher CFG than normal for more interesting results without the prompt and model breaking.

## TL'DR & links

TODO summary
Explanation of why Stable Diffusion is a valuable tool for both beginners and advanced users in the field.

without ui:

```
python scripts/txt2img.py \
 --prompt "Ragnarök in a symbolic and meaningful style, dark_naturalism,
ghostcore, haunted, creepy, trending, award_winning, 8k post-processing,
highly_detailed, dark_fantasy" \
 --n_iter 250 --plms --skip_grid --seed 666 --ddim_steps 100
```

### Links

TODO put where appropriate

- [Stable Diffusion DreamBooth Models](https://cyberes.github.io/stable-diffusion-dreambooth-library/)
- [GFPGAN - Real-world Face Restoration.](https://github.com/TencentARC/GFPGAN)
- [List of science fiction and fantasy artists](https://en.wikipedia.org/wiki/List_of_science_fiction_and_fantasy_artists)

## Full Prompt Examples

- (dr-strange), detailed realistic oil painting of **animals**, trending on artstation, (ultra-detailed:1.3), (masterpiece:1.2), highly detailed, fine detail, realistic proportions, captivating, perfect, intricate, realistic lighting, (rendered by Blender:1.2),
- a forgotten underdark city, (((crypt door))), ((([open |door]2:1))), ((close up)) ((fantasy)) (((dark shadows))) deep underground, black undertones, dramatic, 8k H.R. Giger, ismail inceoglu
- seed: 3605631349
- a hyper detailed realistic full body oil painting of the (queen of blades), sideshow figurines, diablo 4, (lilith), by (DBWayneBarlowe), yusuke murata, hiroya oku, dorian cleavenger, tom bagshaw, (zdzisław beksinski), trending on artstation, (best quality:1.3), (ultra-detailed:1.3), highly detailed face, fine detail, realistic proportions, captivating, perfect, intricate, realistic lighting, [lineart]
- a hyper detailed realistic oil painting of **animals** by moebius, trending on artstation, (ultra-detailed:1.3), (masterpiece:1.2), highly detailed, fine detail, realistic proportions, captivating, perfect, intricate, realistic lighting, (rendered by Blender:1.2)
- an **animals** charging at camera, with **background** in the background, (best quality:1.4), (official art:1.3), (ultra-detailed:1.2), (Game cg:1.6), (masterpiece:1.4), fine detail, realistic proportions, captivating, perfect, intricate, realistic lighting, [lineart], genshin impact, azur lane, dmt art, artgerm
- Kodak portra 4 0 0, wetplate, fisheye, award - winning portrait by britt marling, 1 8 8 0 s kitchen, ghost, picture frames, shining lamps, dust, smoke, 1 8 8 0 s furniture, wallpaper, carpet, books, muted colours, wood, fog, plants, flowers
- Very detailed female, steampunk city taken over by plants in the background, d & d, fantasy, with blue flowers, beautiful face, hyperrealism delicate detailed complex, sophisticated, vibrant colors, volumetric lighting, pop surrealism art by Mark Ryden and Anna Dittmann
- (dr-strange:1.1), Full head and shoulders, beautiful female porcelain sculpture, smooth, all white features on a white background, delicate facial features, white eyes, white lashes, detailed tangled white lillies and lillie leaves on the head, volumetric lighting, (ultra-detailed:1.3), (masterpiece:1.2), highly detailed, fine detail, realistic proportions, captivating, perfect, intricate, realistic lighting, (rendered by Blender:1.2), hdr, high dynamic range
- (concept-art:1.3), An elegant winged fairy in the lord of the rings scenery landscape, looking out at a vast lush valley flowers and homes made of mushrooms, stream, sunrise, god's rays highly detailed, vivid color,
- cinematic lighting, perfect composition, 8 k, gustave dore, derek zabrocki, greg rutkowski, belsinski, octane render
- Landscape photography of fundatura ponorului by marc adamus, An elegant winged fairy in the lord of the rings, scenery landscape, looking out at a vast lush valley, flowers and homes made of mushrooms, stream, sunrise, god's rays highly detailed, vivid color, cinematic lighting, perfect composition, 8 k, gustave dore, derek zabrocki, greg rutkowski, belsinski, octane render, morning, mist, rays of light, beautiful
- masterpiece,highly detailed,game cg,upper body,1boy,solid snake,black headband,middle age,blue eyes,darkbrown hair,much facial hair, camoflage millitary uniform,black eyepatch ,holding assault rifle
- (humorous illustration, hyperrealistic, big depth of field, colors, whimsical cosmic night scenery, 3d octane render, 4k, concept art, hyperdetailed, hyperrealistic, trending on artstation:1.1)
- beautiful detailed glow, depth of field, highly detailed background, dynamic lighting, drama, (Game cg:1.2), (art by Sophie Anderson), (art by Nick Alm)
- Film still of a storm-ravaged offshore supply ship, action sequence from disaster film, outstanding cinematography from blockbuster, 8k HD
- post apocalyptic dystopian speedpaint, cyberpunk desert wasteland, junk everywhere, desert bar with many colorful signs, car parts, cactus, ruined building, small bright lights in the grass, messy lineart sketch, borderlands style art, borderlands concept art, digital painting, messy, bright colors, doodle, unfinished, messy brush strokes, speedpainting, rushed, blocky colors, sharp lines
- a futuristic woman selk'nam, neofuturistic highly detailed, disturbing biohorror saliva, marine animal, (jellyfish face) wlop glossy skin, photorealistic, intricately detailed
- clayitization, character concept of a girl wearing a futuristic grey astronaut suit with orange glowing liquid pipes. coherent face, cute button anime button eyes. vibrant colors, soft lighting, adventurous, atmospheric lighting, 8 k, octane render. by makoto shinkai, stanley artgerm lau, wlop, rossdraws, james jean, andrei riabovitchev, marc simonetti, krenz cushart, sakimichan, d & d trending on artstation, digital art , colorful, intricate details, 3d, 4k, 8k, octane render, Chiaroscuro lighting
- 2d ferocious $ANIMAL head, vector illustration, angry eyes, football team emblem logo, 2d flat, centered, (icon:logo like us football team)
- the One Ring forged in secret in the fires of Mount Doom, which the Dark Lord Sauron poured his cruelty, his malice, and his will to dominate all life
- Aetherpunk style painting of monkeys
- Alien organism, space art, futuristic
- tall rectangular black monolith, monkeys in the desert looking at a large tall monolith
- ultrarealistic cinematic cyborg dragon holding sword in a field of pale pink flowers
- Zombie rock band, at concert
- far away city in dark rock caverns, stalagmites, stone houses, green swirling magic, foggy, dark, mythical, gargoyles, massive ceiling, single white tall concrete tower in the center, massive marble tower in center, green lanterns, purple vegetation
- complex humanoid engine in urban landscape
- skinny bald man with glowing eyes and a long beard at a work bench, hitting a green crystal with a glowing hammer, sparks flying everywhere, a dark lab
- Cyberpunk city oil painting, people in the street, reflection on ground
- itasha in a parking lot, night, neon signs
- Dark Naturalism landscape with trees, Ghostcore, dark fantasy, mystical
- A cavern full of massive magical glowing crystals
- Stone path leading to an abandonned asian temple, asiatic forest
- a persian market, crowded
- (war :1.5) in space, (lasers:1.4), spacestation in space, people
- digital art, game background, dnd, character design, in the style of hearthstone, extreme perspective, ((((fisheye camera))))
- skull with four eyes and symbol on forehead
- battle scene between humans and robots in the future
- dynamic comic hero pose, detailed city at night background
- Oil pastel portrait of a golden retriever with its tongue out, green yellow gradient background
- full page antique lithograph of gemstones minerals
- Demons dance to the chants of the ritual, Black magic wicked Voodoo, digital painting
- Eternal misery when they beg the sire, For the punishments, heart of unholiness
- abstract 3d icon art in cyberpunk style, black background
- abstract digital summit in cybersynth style, octance engine, 8k post-processing highly detailed, cell shading black background
- grinning gothic comic art in the style of hr giger
- female asian warrior, digital art, game background, dnd, character design
- grotesque gothic icon in the style of hr giger, highly detailed, sharp monochrome, dark, award winning
- 3d icon in the style of hr giger, 8k post-processing highly detailed, sharp monochrome, dark, award winning
- Dark Naturalism landscape with dead tree in the style of hr giger, Ghostcore, trending, award winning
- Cat and Dog in love, Storybook, geometric, in the style of Osamu Tezuka, kawaii, award winning
- Artwork by Steve Ditko , highly detailed, award winning
- Science fiction planet by Frank Lewecke, award winning, highly detailed, space art
- The sand beach as grey as a dead cheek, a green tideflow reflects cloud ripples; I stand on the dark wet edge. Cold foam cleanses my toes. I smell driftwood smoke.
- Soundsystem, circle background, cinematic lighting, global illumination, radiant light, crepuscular ray, ((aesthetic, captivating, award winning,HDR, vivid, no dull, hyper-detailed, realistic shaded, fine detail, realistic proportions, symmetrical, sharp focus, 8K resolution, with lineart flat ink))
- (Doom), Ghostcore, Black magic wicked Voodoo, trending, award winning, (((in a symbolic and meaningful style))), High Contrast, Full-HD, AR, Good, Refreshing, Array, Visual Perception, Studio Lighting, De-Noise, Post Processing, (( hyper-detailed and intricate, realistic shaded, fine detail, realistic proportions, symmetrical, sharp focus, 8K resolution, with lineart flat ink, trending on pixiv fanbox))
- ((XTC Art)), in a symbolic and meaningful style, Polaroid, Exposure, Shutter Speed 1/1000, Gamma, Tones of Black in Background, High Contrast, Full-HD, AR, Good, Refreshing, Array, Visual Perception, Studio Lighting, De-Noise, Post Processing
- The death of God, aesthetic, captivating, award winning, technicolor, HDR, vivid, stunning color palette, no dull, concept art, hyper-detailed and intricate, realistic shaded, fine detail, realistic proportions, symmetrical, sharp focus, 8K resolution, with lineart flat ink, trending on pixiv fanbox
- portrait of a model by Flora Borsi, style by Flora Borsi, bold, bright colours, Mohawk haircut, ((Flora Borsi))
- ethereal "very close face shot / close portrait photo" of pixie fairy **\_\_\_\_** as a (magical fairy) wearing green ((delicate leave skirt and top made of leaves)) (flower head dress) ((ancient lush woods)) colourful mystical vibrant enchanting, cinematic lighting, global illumination, depth of field, sub surface scattering, ray tracing, 8k photo
- Retro comic style artwork, highly detailed "James Bond", comic book cover, symmetrical, vibrant
- ferocious red hannya mask, highly detailed, with gold and red shield, magical mist around it, illustration, angry eyes, 2d flat, centered
- Hologram of a XY floating in space, a vibrant digital illustration, dribbble, quantum wavetracing, black background, behance hd
- portrait epic fantasy art, evil black (knight), (face portrait), detailed (helmet closeup), (two red smoldering eye sockets:1.5) looking at camera, symmetrical centered head and shoulders portrait, (two|Viking horns on helmet), (two red eyes) inside horned barbute helmet, (( two eyes on face)), good anatomy, [color grading], correct anatomy, dark atmosphere, saturated complementary colors, awesome, smooth black metal armor, spiked shoulders,(shadowed|hidden face) inside barbute visor, sunset background, from front view, haunted knight

---

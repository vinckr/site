---
layout: layouts/page.njk
tags: blog
title: Stable Diffusion Guide
---

## Introduction

This guide is intended for people new to image generation using Stable Diffusion (SD).  
The main technology used in this guide is SD 1.6 running interference on a local computer. Interference means in our case generating images from a text prompt or another image.

In this guide you will learn the basics of how to run a SD model on your local computer. Most of the prompting techniques and workflows from this guide are also applicable to other models, but SD 1.6 is the main base model used in this guide.

### What is Stable Diffusion

- [Stable Diffusion explained with visuaotherlization](https://poloclub.github.io/diffusion-explainer/)
- [Stable Diffusion on Wikipedia](https://en.wikipedia.org/wiki/Stable_Diffusion)
- [A walk through latent space with Stable Diffusion](https://keras.io/examples/generative/random_walks_with_stable_diffusion/)
- [The Huggingface Stable Diffusion Guide ](https://huggingface.co/docs/diffusers/v0.14.0/en/stable_diffusion)

### Why SD 1.6 if there is a newer version

Stable Diffusion 1.6 has a huge number of trained models, so you have a huge variety of different styles to choose from.

Stable Diffusion 2.x is not great for a couple of reasons, to be honest I have not used it much since it was a bit of a hassle to set up.

Stable Diffusion XL is great but has in my experience less variety than 1.6. I would recommend to check out SDXL, you can get some great results in higher resolution with it.

## Getting Started

<!-- TODO rewrite & investigate how much RAM you need and so on -->

To run SD on your computer you need a good amount of GPU RAM. It works well with 8-16 gb of RAM, but it also works on much weaker machines - it just takes much longer to generate.

- Guide how to run Stable Diffusion on a M1 Macbook without user interface: [Run Stable Diffusion on your M1 Mac's GPU](https://replicate.com/blog/run-stable-diffusion-on-m1-mac)
- Guide how to run Stable Diffusion on a Mac with user interface: [Stable Diffusion web UI repo on GitHub](https://github.com/AUTOMATIC1111/stable-diffusion-webui)

For beginners I recommend to use the AUTOMATIC1111 Stable Diffusion web UI. Its a convenient interface for Stable Diffusion and has many features and extensions.  
You can go from generating images from a simple prompt to more elaborate and complex setups without writing any code. [Overview of Stable Diffusion web UI features](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Features)

I am still catching up and e week there is something new to play around with. There are many other good UIs for running Stable Diffusion, the most notable would be [ComfyUI](https://github.com/comfyanonymous/ComfyUI). Comfy is node-based and more complex than the web UI. I would recommend Comfy for experienced users. Comfy gives you much more control over the whole process but also requires some more understanding of how it all works.

### Installation

- [Install and Run on NVidia GPUs](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Install-and-Run-on-NVidia-GPUs)
- [Installation on Apple Silicon](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Installation-on-Apple-Silicon)
- [Install and Run on AMD GPUs](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Install-and-Run-on-AMD-GPUs)

## Terminology

Some common terms and their meaning in the context of Stable Diffusion. If you are interested in the details and technical background, please refer to the resources in [Links](###Links).

### Model

The model in Stable Diffusion refers to the entire system used for generating images from text prompts. It consists of several key components, including a variational autoencoder (VAE), forward and reverse diffusion processes, a noise predictor, and text conditioning.

### Latent Space

The latent space refers to the high-dimensional vector space in which the model encodes the input data. In the context of Stable Diffusion, the latent space is where the model encodes the images during the generation process. The model starts from a noisy initial image in this latent space and gradually refines - denoises - it in a series of steps to produce the final image.

### Token

A token refers to the smallest unit of data that the model can understand and process. Tokens are typically words or subwords derived from the input text prompts. The text prompts are tokenized using a process like the one implemented by the [CLIP tokenizer](https://huggingface.co/docs/transformers/model_doc/clip), which breaks down each word in a textual prompt and embeds this data into a vector.

### Steps

Steps refer to the iterations in the process of either generating an image - also called inference steps - or training the model. Steps most often refer to the number of denoising steps used in the generation process. These steps represent iterations of noise removal from a noisy initial image, gradually refining it to produce the final generated image. The number of steps can be a key hyperparameter that users can control to influence the generation proces. Usually you get good results using a low number of steps, so when in doubt reduce the number of steps and see how it influences your generation. Reducing the number of steps also makes the generation process faster.

## Text to Image

- [Learn how to use Text to Image, basic prompting and negative prompts here](./txt2img.md)

## Image to Image

With img2img instead of starting out initially with just random noise you use take another image. Prompting img2img works the same.

I use img2img for:

- Refining an image generated through text.
- Using an existing image as the basis for a new generation.

No guide [here](./img2img.md) yet.

## Inpainting

No guide [here](./inpainting.md) yet.

## Outpainting

No guide [here](./outpainting.md) yet.

## Advanced Usage of Stable Diffusion

It gets complex really fast, so hold onto your horses.

You can train new models on other images, to refine the base Stable Diffusion model. Think of it as Stable Diffusion is trained on all images in the internet, also the really bad quality ones. Now people create a new checkpoint (like a 7gb model file that you download) with for example high quality images. The resulting model will have a strong bias towards those higher quality images, resulting in better quality outputs. Now depending on how the model is trained this can be a stronger or weaker effect.
The real magic happens when there are dozens of high quality models trained and then combined in dozens of layers, so you have a model like

### Model improvements

There are multiple ways to improve the outputs from the core Stable Diffusion models. The most common ones are checkpoints, loras, embeddings, and hypernetworks.

Checkpoints are powerful and easy to use.
Loras are almost as powerful as checkpoints and much more lightweigth - Loras are the best option for refining your generations!
Embeddings need a lot of care to produce good results.
Hypernetworks are similar to embeddings but easier to use.

A brief explanation of each technique follows. If you want to learn more refer to the [Links](#links) section. In most cases you don't need to understand the technical background to use a technique

#### Checkpoints

A checkpoint model is similar to the base Stable Diffusion model [described earlier](#model). The weights in the model are adjusted using training data to generate images resembling the training data. Training a new model requires caution to prevent overwriting existing data or the rendered images will resemble only the training data. The checkpoint model has a large filesize, approximately 4-7GB.

<!-- todo checkpoint examples-->

##### VAE

The VAE (Variational Autoencoder) is a component that is sometimes included with checkpoint models. In the web UI, you can select the VAE that came with your checkpoint model or leave it on "Auto" for best results in most cases. If your images appear dull, lackluster, or have strange small colored dots or circles in them, it may indicate that you are using the incorrect VAE. This is normally not a concern, but I mention it because it caused confusion for me and I spent weeks generating images with the wrong VAE before understanding the issue.

<!-- todo wrong vae examples-->

#### Loras

LoRA or Low-Rank Adaptation is a training method for refining Stable Diffusion models. It helps adapt a pre-trained model to a specific task or dataset efficiently, without the need to train a new model from the beginning, which can be expensive and time-consuming. By adjusting the weight of LoRA during the image generation process, you can control its impact. A higher weight gives LoRA more influence, while a lower weight reduces its influence. Typically, a LoRA has a moderate file size, usually less than 1 GB. I would recommend to use Loras whenever possible because they are simpler to train than checkpoints, much smaller filesize, and have a big impact on the output.

<!-- todo lora examples-->

#### Embeddings

An embedding is created through textual inversion. Textual inversion means a specific prompt is found to generate images similar to the model's training data. The generated content is limited to what the model can already produce. An embedding is essentially a concise and precise prompt for the model and has a small file size of approximately 150MB.

An embedding is basically a shorthand for a long and complex prompt.

<!-- todo embeddings examples-->

#### Hypernetworks

A hypernetwork is an additional layer that modifies the results of a model after the image has been processed. The hypernetwork will skew all results from the model towards your training data. This means that if an image contains something related to the training data, it will resemble the training data. For example If you trained a hypernetwork on a specific cat, you will have a hard time trying to get any other cat using that hypernetwork. Hypernetworks rely on keywords known to the model and have a small file size of approximately 80mb.

<!-- todo hypernetworks examples-->

## Advanced techniques

### Weights

- weights influence how much the model steers the output towards that token
- () around a token increases the weight of that token, for example

By default each token has a weight of 1. If you put parantheses around it the weight will increase to 1.1, double parantheses to 1.2 etc. You can also specify the weight directly, for example (example-token:1.4) has a weight of 1.4.

In my experience it makes sense to emphasize a few tokens that are important for your generation, but don't overdo it. It got the best results by keeping the weights between 1 and 1.5.

<!-- todo examples-->

### Mixing prompts

- Mix prompts, for example [duck|pig].

No guide here yet.

<!-- todo examples-->

### X/Y inputs

No guide here yet.

<!-- todo examples-->

### Wildcards

No guide here yet.

<!-- todo examples-->

### ControlNet

No guide here yet.

<!-- todo examples-->

### Upscaling

No guide here yet.

## Workflow examples

#### Example img2img workflow

- Generate many picture with your prompt and 512x512 pixels.
- Once you have a good image, use it as a base sketch for img2img.
- Increase the resolution while setting the denoising to approximately 0.60. Afterwards, repeatedly process the image using a denoising value of 0.2-0.3.
- Use an image editor to refine the image
  - To remove any bad looking elements, use a brush of the same color as the adjacent detail.
  - To add refinements, like enhancing fingers, use a pencil with a contrasting color, such as black.
  - Always reprocess the edited image with a denoising value of 0.3.
- When content with the result, upscale the image to a higher resolution.

- General workflow: Generate a ton of images, look for cool poses, send to inpainting. Work on hands and bad anatomy with mask blur 4, inpaint at full resolution, masked content: original, 32 padding, denoise 0.2-0.4 for small changes, 0.55-0.75 for large changes. Once you have anatomy and hands nailed down, move on to cosmetic changes to clothing, then faces. I cranked up the number of steps for faces, no idea if that actually helps.

### Example inpainting workflow

- Generate many images, select interesting results, and send to inpainting.
- Work on hands, weird spots, and bad anatomy with mask blur set to 4.
- Inpaint the images at their original resolution while keeping the masked content intact. Use a padding of 32.
- For minor changes, apply a denoise value between 0.2 and 0.4. For major changes, use a denoise value between 0.55 and 0.75.
- Rinse and repeat for other elements of the generation.

## Training your own

### Dreambooth models

- [What is Dreambooth](https://dreambooth.github.io/)

> Disclaimer: I did not test this method yet, as using and training Loras is now much better and easier than Dreambooth.

- Provide a minimum of 20 close portrait shots, 4 middle distance torso shots, 4 long distance full length body shots, and 1 seated shot.
- Use clear, high-resolution images.
- Capture variety by using different environments, lighting, hair styles, expressions, poses, angles, and distances from the subject.
- Remove images with hands too close to or touching the head to avoid unwanted results.
- Exclude selfie images to avoid the fisheye lens effect.
- Ensure images have no visible cropped borders to prevent white borders in the renders.
- Avoid fake heavy depth of field or bokeh in the images to maintain a natural appearance.
- Edited the face in training images to address issues such as artifacts or model breakdown.
- Class name should be "a **\_**". Manually generate class images by adding an "a" before the class name to increase rendering accuracy. This helps avoid producing unrelated results. Additionally, use 60 training images, 2000 "person" class images, and train for 6000 steps.
- A well-trained model reduces efforts for Stable Diffusion and allows for higher CFG for more interesting results without disrupting the prompt or model performance.

## Links

### Stable Diffusion Explained

- [The Illustrated Stable Diffusion](https://jalammar.github.io/illustrated-stable-diffusion/)
- [Stable Diffusion explained with visualization](https://poloclub.github.io/diffusion-explainer/)
- [Stable Diffusion on Wikipedia](https://en.wikipedia.org/wiki/Stable_Diffusion)
- [A walk through latent space with Stable Diffusion](https://keras.io/examples/generative/random_walks_with_stable_diffusion/)
- [The Huggingface Stable Diffusion Guide](https://huggingface.co/docs/diffusers/v0.14.0/en/stable_diffusion)
- [What is a Variational Autoencoder (VAE)?](https://jaan.io/what-is-variational-autoencoder-vae-tutorial/)

### Stable Diffusion Resources

- [Stable Diffusion WebUI](https://github.com/AUTOMATIC1111/stable-diffusion-webui)
  - [Installation](https://github.com/AUTOMATIC1111/stable-diffusion-webui#installation-and-running)
  - [Features](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Features)
  - [Extensions](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/)
- [Download Models, Embeddings, Loras, etc.](https://civitai.com/)
- [Stable Diffusion beginner's guide](https://stable-diffusion-art.com/beginners-guide/)
- [Video to video with Stable Diffusion](https://stable-diffusion-art.com/video-to-video/)
- [ControlNet guide](https://stable-diffusion-art.com/controlnet/)
- [Search other peoples prompts](https://search.krea.ai/)
- [Stable Diffusion DreamBooth Models](https://cyberes.github.io/stable-diffusion-dreambooth-library/)
- [GFPGAN - Real-world Face Restoration.](https://github.com/TencentARC/GFPGAN)
- [List of science fiction and fantasy artists](https://en.wikipedia.org/wiki/List_of_science_fiction_and_fantasy_artists)
- [Stable Diffusion LoRA Guide (Downloading, Usage, Training)](https://aituts.com/stable-diffusion-lora/)

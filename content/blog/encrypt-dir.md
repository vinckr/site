---
layout: page.njk
tags: blog
title: How to encrypt files in a git repository easy mode
---

In this guide you will learn about a solution to the problem of how to commit drafts and notes to a public git repository, while not making them readable for others. You can also use this with any kind of files other than text; images, code, videos, you name it.

## Why

I want to open source my website source code, so I can share the code with people and let them roast my CSS. One problem: I don't want to share my draft texts publicly, but I would still like to include them in the git repository. This way I don't have to rely on another service to store draft text and they get versioned with the rest of the website.

It would cause me some anxiety if I always had in the back of my mind that everyone could read my rough drafts. At the same time I enjoy going through other peoples personal website stacks and seeing all the creative, cute, and genius things people build in their personal site. I dont think my website is all that special, it still feels good to put it out there. You can now find the source code for my website on [GitHub](https://github.com/vinckr/vinckr.com).

My solution is to use a command-line tool to encrypt my drafts and add the encrypted files whenever I make a commit. I can `git pull` them together with the full website stack in one repository from any machine and handle everything with `make` commands.

## Lets go encrypt-dir

[`encrypt-dir`](https://github.com/ory/encrypt-dir) is a command-line utility that allows you to encrypt directories through the CLI. The tool uses the AES-GCM encryption algorithm with a 128-bit key to encrypt data.

The key feature of the `encrypt-dir` tool is the ability to encrypt files in a user-specified directory. This means you can encrypt all the files in a particular folder with one command. You specify a keyphrase for the encryption.

Note that there is some information leaked through the filename, as that's still publicly readable.

### Install

If you don't already have Go installed, download, and install it from the [Go website](https://golang.org/). The use `go get` to download and install `encrypt-dir`.

```
go get github.com/ory/encrypt-dir
```

After `encrypt-dir` has been installed, you can verify that it's working by running the `encrypt-dir --help` command in your terminal or command prompt.
To use `encrypt-dir` to encrypt files, you will need to specify the directory containing the files and the password for the encryption. You can do this using the `-d` flag for directory and `-p` for password. For example, to encrypt all files in the `_drafts` directory using the password mypassword, you would run the following command:

```
encrypt-dir -d \_drafts -p mypassword
```

This will encrypt all files in the `_drafts` directory with `mypassword` and save the result `drafts.secure`.

## Makefile example

The following is an example of how I use `encrypt-dir` in a Makefile.

First, to make `encrypt-dir` available you need to build it:

```
.bin/encrypt-dir:
	echo "Building encrypt-dir"
	go build -o .bin/encrypt-dir github.com/ory/encrypt-dir
```

To encrypt the drafts, run `make encrypt-drafts`. This zips the contents of the `_drafts` folder, saves them in `drafts` and encrypts them.
I also added `_drafts` and `drafts/drafts.zip` to gitignore to avoid getting them into git unencrypted. The encryption key is just saved locally as well of course.

```
encrypt-drafts: .bin/encrypt-dir # encrypt files in \_drafts folder
echo "Encrypting drafts"
zip -r drafts/drafts.zip \_drafts/\*
echo "${ENCRYPTION_KEY}"
	.bin/encrypt-dir encrypt --key=${ENCRYPTION_KEY} drafts
```

To decrypt, just run `make decrypt-drafts`. I need this when I have a fresh clone of the repository.

```
decrypt-drafts: .bin/encrypt-dir # decrypt files in drafts folder
echo "Decrypting drafts"
.bin/encrypt-dir decrypt --key=${ENCRYPTION_KEY} drafts
unzip -o drafts/drafts.zip
```

To use the `encrypt-dir` tool to encrypt files in your GitHub repository, you clone your repository, run `make encrypt-drafts` to encrypt the files, commit the changes, and push the changes to your remote repository on GitHub.

## Closing thoughts

I have been using this method for a some months now, and to me it's convenient. `encrypt-dir` is a versatile tool, that can also be used to encrypt secrets, API keys or similar and run it on your server. I'm sure there are probably other/smarter/simpler ways to encrypt your drafts, but setting up and learning about the tool was fun. You can learn how to use it in minutes.

In case you missed it you can find `encrypt-dir` on [GitHub](https://github.com/ory/encrypt-dir) or [pkg.go.dev](https://pkg.go.dev/github.com/ory/encrypt-dir/cmd).

~Do you know a better method or how to do this in a bash-oneliner? Let me know!~

Right after completing this guide, GPT tells me I can do the same thing with this bash one-liner... Well maybe some other time 🫠

```
zip -r - ./_drafts | openssl aes-256-cbc -pbkdf2 -salt -out files.zip.encrypted
```

Thanks for reading 👋

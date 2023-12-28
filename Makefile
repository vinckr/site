ifneq ("$(wildcard .env)","")
    include .env
endif

help:  # show all available Make commands
	@cat Makefile | grep '^[^ ]*:' | grep -v '^\.bin/' | grep -v '^node_modules' | grep -v '.SILENT:' | grep -v help | sed 's/:.*#/#/' | column -s "#" -t
	
dev: # run a local server to preview the site
	npm run start

format: .bin/shfmt node_modules  # format the source code
	echo "\n formatting ..."
	.bin/shfmt --write .
	npm exec -- prettier --write .

links: # generate markdown links from urls.txt
	echo "Generating markdown links"
	go run cmd/links/main.go &> urls-sorted-markdown.txt

links-sort: # sort links and delete duplicates
	echo "Sorting markdown links"
	cat urls-unsorted.txt | sort -u > urls-sorted.txt

encrypt-drafts: .bin/encrypt-dir # encrypt files in _drafts folder
	echo "Encrypting drafts"
	zip -r drafts/drafts.zip _drafts/*
	echo "${ENCRYPTION_KEY}"
	.bin/encrypt-dir encrypt --key=${ENCRYPTION_KEY} drafts
	git add .
	git commit -m "chore: drafts"

decrypt-drafts: .bin/encrypt-dir # decrypt files in drafts folder
	echo "Decrypting drafts"
	.bin/encrypt-dir decrypt --key=${ENCRYPTION_KEY} drafts
	unzip -o drafts/drafts.zip

check-links: # check links in markdown files
	echo "Checking links"
	cd content && find . -name \*.md -print0 | xargs -0 -n1 markdown-link-check

.bin/encrypt-dir:
	echo "Building encrypt-dir"
	go build -o .bin/encrypt-dir github.com/ory/encrypt-dir 

node_modules: package.json package-lock.json
	echo installing Node dependencies ...
	npm ci
	touch node_modules  # update timestamp so that Make doesn't reinstall it over and over

.SILENT:
.DEFAULT_GOAL := help

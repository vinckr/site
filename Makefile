ifneq ("$(wildcard .env)","")
    include .env
endif

u: # format, encrypt drafts, and commit to git
	make format
	make test
	make encrypt-drafts
	git add .
	git commit -m "chore: format + drafts"

help:  # show all available Make commands
	@cat Makefile | grep '^[^ ]*:' | grep -v '^\.bin/' | grep -v '^node_modules' | grep -v '.SILENT:' | grep -v help | sed 's/:.*#/#/' | column -s "#" -t

build: .bin/gokesh # build HTML without committing
	.bin/gokesh page index
	.bin/gokesh page about
	.bin/gokesh dir blog
	
dev: # run a local server to preview the site
	make build
	@echo "Preview running at http://localhost:8000"
	go run cmd/dev/main.go

build: .bin/gokesh # build HTML without committing
	.bin/gokesh page index
	.bin/gokesh page about
	.bin/gokesh dir blog
	
dev: # run a local server to preview the site
	make build
	@echo "Preview running at http://localhost:8000"
	go run cmd/dev/main.go

format: .bin/shfmt node_modules  # format the source code
	echo "\n formatting ..."
	.bin/shfmt --write .
	npm exec -- prettier --write .

test: .bin/shellcheck .bin/shfmt node_modules  # run all linters
	echo running tests ...
	find . -name '*.sh' | xargs .bin/shellcheck
	echo Verifying formatting ...
	.bin/shfmt --list .

links: # generate markdown syntax links from urls.txt
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

decrypt-drafts: .bin/encrypt-dir # decrypt files in drafts folder
	echo "Decrypting drafts"
	.bin/encrypt-dir decrypt --key=${ENCRYPTION_KEY} drafts
	unzip -o drafts/drafts.zip

.bin/gokesh:
	echo "Building gokesh"
	go mod download
	go build -o .bin/gokesh github.com/vinckr/gokesh/cmd/build

.bin/encrypt-dir:
	echo "Building encrypt-dir"
	go build -o .bin/encrypt-dir github.com/ory/encrypt-dir 

.bin/shellcheck: Makefile
	echo "\n installing Shellcheck ..." 
	curl -sSL https://github.com/koalaman/shellcheck/releases/download/stable/shellcheck-stable.linux.x86_64.tar.xz | tar xJ
	mkdir -p .bin
	mv shellcheck-stable/shellcheck .bin
	rm -rf shellcheck-stable
	touch .bin/shellcheck   # update the timestamp so that Make doesn't re-install the file over and over again

.bin/shfmt: Makefile
	echo installing Shellfmt ...
	mkdir -p .bin
	curl -sSL https://github.com/mvdan/sh/releases/download/v3.6.0/shfmt_v3.6.0_darwin_arm64 -o .bin/shfmt
	chmod +x .bin/shfmt

node_modules: package.json package-lock.json
	echo installing Node dependencies ...
	npm ci
	touch node_modules  # update timestamp so that Make doesn't reinstall it over and over

.SILENT:
.DEFAULT_GOAL := help

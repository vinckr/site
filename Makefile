include .env

.bin/encrypt-dir:
	echo "Building encrypt-dir"
	go build -o .bin/encrypt-dir github.com/ory/encrypt-dir 

# update, builds HTML pages and commits all changes to git
u:
	make build
	make encrypt-drafts
	git add .
	git commit -m "chore: update"

# build HTML without committing
build:
	echo "Building HTML files"
	go run cmd/build/main.go

# sort links and delete duplicates
links-sort:
	echo "Sorting markdown links"
	cat urls-unsorted.txt | sort -u > urls.txt

# generate markdown syntax links from urls.txt
links:
	echo "Generating markdown links"
	go run cmd/links/main.go

# encrypt all files in _drafts folder
encrypt-drafts: .bin/encrypt-dir
	echo "Encrypting drafts"
	zip -r drafts/drafts.zip _drafts/*
	echo "${ENCRYPTION_KEY}"
	.bin/encrypt-dir encrypt --key=${ENCRYPTION_KEY} drafts

# decrypt all files in drafts folder and unzip in _drafts
decrypt-drafts: .bin/encrypt-dir
	echo "Decrypting drafts"
	.bin/encrypt-dir decrypt --key=${ENCRYPTION_KEY} drafts
	unzip drafts/drafts.zip -d _drafts

# decrypt:
#  encrypt-dir decrypt -k $(shell cat .env | grep ENCRYPTION_KEY | cut -d '=' -f2) $(shell cat .env | grep DECRYPT_DIR | cut -d '=' -f2)

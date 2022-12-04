
build:
	echo "Building HTML files"
	go run cmd/build/main.go

links-sort:
	echo "Sorting markdown links"
	cat urls-unsorted.txt | sort -u > urls.txt

links:
	echo "Generating markdown links"
	go run cmd/links/main.go

u:
	make build
	git add .
	git commit -m "chore: update"




build:
	echo "Building HTML files"
	go run cmd/build/main.go

links:
	echo "Generating markdown links"
	go run cmd/links/main.go

u:
	make build
	git add .
	git commit -m "chore: update"



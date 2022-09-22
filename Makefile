
build:
	echo "Building HTML files"
	go run build.go

u:
	make build
	git add .
	git commit -m "chore: update"



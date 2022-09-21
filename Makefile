u:
	echo "update"
	git add .
	git commit -m "chore: update"


build:
	echo "Building HTML files"
	go run build.go
	git add .
	git commit -m "chore: lazy update"



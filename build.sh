#!/bin/sh
set -e

echo "Install build tools"
go get github.com/vinckr/gokesh/cmd/build
go install github.com/vinckr/gokesh/cmd/build
echo "Building HTML files"
go run vinckr/gokesh/cmd/build/ page index
go run vinckr/gokesh/cmd/build/ page about
go run vinckr/gokesh/cmd/build/ dir markdown/blog/

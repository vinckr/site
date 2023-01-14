#!/bin/sh
set -e

echo "Install build tools"
go build -o .bin/build github.com/vinckr/gokesh/
echo "Building HTML files"
.bin/build go run vinckr/gokesh/cmd/build/ page index
bin/build go run vinckr/gokesh/cmd/build/ page about
bin/build go run vinckr/gokesh/cmd/build/ dir markdown/blog/

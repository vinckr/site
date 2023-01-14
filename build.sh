#!/bin/sh
set -e

cd .bin/build 
echo "Building tools"
go build -o .bin/build github.com/vinckr/gokesh/
echo "Building HTML files"
.bin/build go run vinckr/gokesh/cmd/build/ page index
.bin/build go run vinckr/gokesh/cmd/build/ page about
.bin/build go run vinckr/gokesh/cmd/build/ dir markdown/blog/

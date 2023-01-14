#!/bin/sh
set -e

echo "Create build and output directory"
mkdir -p .bin/build
mkdir -p .bin/build/public
echo "Copying files"
cp -R markdown .bin/build
cp -R templates .bin/build
echo "Installing tools"
cd .bin/build
go get github.com/vinckr/gokesh/cmd/build
echo "Building tools"
go build -o .bin/build github.com/vinckr/gokesh/cmd/build
echo "Building HTML files"
.bin/build page index
.bin/build page about
.bin/build dir markdown/blog/

echo "Copying public files"
cp -R public/* ../../public

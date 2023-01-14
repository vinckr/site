#!/bin/sh
set -e

echo "Building HTML files"
go run vinckr/gokesh/cmd/build/ page index
go run vinckr/gokesh/cmd/build/ page about
go run vinckr/gokesh/cmd/build/ dir markdown/blog/

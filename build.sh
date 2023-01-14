#!/bin/sh
set -e


echo "Install build tools"
mkdir -p functions
GOOS=linux
GOARCH=amd64
GO111MODULE=on
GOBIN=${PWD}/functions go get ./...
GOBIN=${PWD}/functions go install ./...
echo "Building HTML files"
go run vinckr/gokesh/cmd/build/ page index
go run vinckr/gokesh/cmd/build/ page about
go run vinckr/gokesh/cmd/build/ dir markdown/blog/

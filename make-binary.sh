#!/bin/bash
# Script to build umotd binary for Linux on x86_64 and arm64

VERSION=0.1.3

# x86_64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o umotd_"${VERSION}"_linux_amd64

# arm64
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o umotd_"${VERSION}"_linux_arm64

echo "Built umotd v$VERSION for Linux on x86_64 and arm64"

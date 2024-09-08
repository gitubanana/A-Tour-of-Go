#!/usr/bin/env sh

if [ -n "$(find . -type f -name "*.go" -exec gofmt -l -w -d {} \;)" ]; then
    echo "Not formatted"
    exit 1
else
    echo "All formatted"
fi
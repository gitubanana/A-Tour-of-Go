#!/usr/bin/env sh

find . -type f -name "*.go" -print -exec go fmt {} \;
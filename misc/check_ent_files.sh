#!/usr/bin/env bash
go generate ./...
status=$(git status --porcelain)
if [ -n "$status" ]; then
    echo "you need to run 'go generate ./...' and commit the changes"
    echo "$status"
    exit 1
fi

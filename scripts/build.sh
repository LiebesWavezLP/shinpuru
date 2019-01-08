#!/bin/bash

TAG=$(git describe --tags)
if [ "$TAG" == "" ]; then
    TAG="untagged"
fi

COMMIT=$(git rev-parse HEAD)

echo "Getting dependencies..."
go get -v -t ./...

echo "Building..."
go build -ldflags " \
    -X github.com/zekroTJA/shinpuru/util.AppVersion=$TAG \
    -X github.com/zekroTJA/shinpuru/util.AppCommit=$COMMIT \
    -X github.com/zekroTJA/shinpuru/util.Release=TRUE"

wait
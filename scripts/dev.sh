#!/bin/bash

go get github.com/go-delve/delve/cmd/dlv

go build -gcflags="all=-N -l" -o $PWD "$PWD/main.go"

echo "Build succeeded!!!!"

dlv --listen=:40000 --headless=true --api-version=2 --accept-multiclient exec "$PWD/main"

# go run -race $PWD
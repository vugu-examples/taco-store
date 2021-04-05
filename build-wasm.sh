#!/bin/bash
go generate ./ui
GOARCH=wasm GOOS=js go build -o dist/index.wasm ./wasm/index
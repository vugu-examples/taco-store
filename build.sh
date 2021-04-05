#!/bin/bash
sh ./build-wasm.sh
go build -o bin/server ./cmd/server
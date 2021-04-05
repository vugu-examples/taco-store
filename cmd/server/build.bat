go generate ../../ui
export GOARCH=wasm export GOOS=js
go build -o ../../dist/index.wasm ../../wasm/index
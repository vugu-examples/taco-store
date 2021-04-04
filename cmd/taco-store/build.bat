go generate ../../ui
export GOARCH=wasm export GOOS=js
go build -o ../../dist/main.wasm ../../wasm/taco-store
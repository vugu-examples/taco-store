go generate ../../ui
GOARCH=wasm GOOS=js go build -o ../../dist/main.wasm ../../wasm/taco-store
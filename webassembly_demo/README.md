cd ./webassembly_demo/cmd/wasm
GOOS=js GOARCH=wasm go build -o ../../assets/json.wasm

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./webassembly_demo/assets/

cd ./webassembly_demo/cmd/server/
go run main.go
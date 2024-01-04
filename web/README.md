# Compiling
To build the web interface from source, you will need to compile the Go code into a WebAssembly file:
```
GOOS=js GOARCH=wasm go build -ldflags="-s -w" index.go
```
This will create a binary file. Compress it with [LZMA](https://github.com/LZMA-JS/LZMA-JS), encode it in Base64, and paste the final result to [L198](https://github.com/HACKERALERT/Picocrypt/blob/main/web/index.html#L198).

You'll also need to update [`wasm_exec.js`](https://cdn.jsdelivr.net/gh/golang/go@go1.21.5/misc/wasm/wasm_exec.min.js) (replace the Go version accordingly) on [L197](https://github.com/HACKERALERT/Picocrypt/blob/main/web/index.html#L197) to glue everything together.

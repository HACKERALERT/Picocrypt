# Compiling
To build the web interface from source, you will need to compile the Go code into a WebAssembly file:
```
GOOS=js GOARCH=wasm go build -ldflags="-s -w" index.go
```
This will create a binary file. Compress it with [LZMA](https://github.com/LZMA-JS/LZMA-JS), encode it in Base64, and paste the final result to [L198](https://github.com/HACKERALERT/Picocrypt/blob/main/web/index.html#L198).

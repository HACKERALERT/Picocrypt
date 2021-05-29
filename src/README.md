# Running From Source
If you would like to run Picocrypt from source, or an executable isn't available for your platform, you've come to the right place. Running from source is very simple, and I've made it very straightforward by writing specific instructions for each platform.

# Windows
1. Install the Go Programming Language from <a href="https://golang.org/dl">here</a>.
2. Download the source file `Picocrypt.go` from above.
3. Install Picocrypt's dependencies:
```cmd
go get -u -v github.com/pkg/browser && go get -u -v github.com/zeebo/blake3 && go get -u -v golang.org/x/crypto/sha3 && go get -u -v golang.org/x/crypto/argon2 && go get -u -v github.com/AllenDang/giu && go get -u -v github.com/OpenDiablo2/dialog && go get -u -v golang.org/x/crypto/blake2b && go get -u -v golang.org/x/crypto/blake2s && go get -u -v github.com/atotto/clipboard && go get -u -v github.com/klauspost/reedsolomon && go get -u -v golang.org/x/crypto/chacha20poly1305 && go get -u -v github.com/HACKERALERT/Picocypher/monocypher
```
4. Open a Command Prompt in the directory which contains `Picocrypt.go` and build Picocrypt from source:
```cmd
go mod init Picocrypt && go mod tidy && go build -ldflags "-s -w -H=windowsgui -extldflags=-static" Picocrypt.go
```
5. You should now see `Picocrypt.exe` in your directory. Now, go and download the `sdelete64.exe` from the list above and place it in the same directory as `Picocrypt.exe`.
6. You are now complete and you can double click `Picocrypt.exe` to run Picocrypt.

# macOS
Coming soon....

# Linux
Coming soon....

## Other
If your distro is not Debian-based, don't worry! Building from source is still very simple.
1. Install the Go Programming Language from <a href="https://golang.org/dl">here</a>.
2. Clone this repository using `git clone` or by downloading the master branch.
3. Open a terminal where `Picocrypt.go` is located (in the `src` directory).
4. Get Picocrypt's dependencies:
```bash
go get -u -v github.com/pkg/browser && go get -u -v github.com/zeebo/blake3 && go get -u -v golang.org/x/crypto/sha3 && go get -u -v golang.org/x/crypto/argon2 && go get -u -v github.com/AllenDang/giu && go get -u -v github.com/OpenDiablo2/dialog && go get -u -v golang.org/x/crypto/blake2b && go get -u -v golang.org/x/crypto/blake2s && go get -u -v github.com/atotto/clipboard && go get -u -v github.com/klauspost/reedsolomon && go get -u -v golang.org/x/crypto/chacha20poly1305 && go get -u -v github.com/HACKERALERT/Picocypher/monocypher
```
5. Build from source:
```bash
go mod init Picocrypt && go mod tidy && go build -ldflags "-s -w" Picocrypt.go
```
6. You're all done. You can now run the file `Picocrypt`.

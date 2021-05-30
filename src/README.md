# Running From Source
If you would like to run Picocrypt from source, or an executable isn't available for your platform, you've come to the right place. Running from source is very simple, and I've made it even easier with these simple instructions. Note that the instructions are generic and will work on any platform.

1. Prerequisites (only for Linux):
```bash
sudo apt-get install -y gcc make curl git tar wget xz-utils libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libgl1-mesa-dev libxxf86vm-dev libgtk-3-dev xdg-utils && sudo apt-get install -y libglx-dev || echo "" && sudo apt-get install -y xclip || sudo apt-get install xsel
```
2. Install the Go Programming Language from <a href="https://golang.org/dl">here</a>, or from your package manager. Go 1.16 or higher is recommended.
3. Download the source file `Picocrypt.go` from above or `git clone` this repository.
4. Install Picocrypt's Go dependencies:
```bash
go get -u -v github.com/pkg/browser && go get -u -v github.com/zeebo/blake3 && go get -u -v golang.org/x/crypto/sha3 && go get -u -v golang.org/x/crypto/argon2 && go get -u -v github.com/AllenDang/giu && go get -u -v github.com/OpenDiablo2/dialog && go get -u -v golang.org/x/crypto/blake2b && go get -u -v golang.org/x/crypto/blake2s && go get -u -v github.com/atotto/clipboard && go get -u -v github.com/klauspost/reedsolomon && go get -u -v golang.org/x/crypto/chacha20poly1305 && go get -u -v github.com/HACKERALERT/Picocypher/monocypher
```
5. Initialize Go and tidy the modules:
```bash
go mod init Picocrypt && go mod tidy
```
6. Go to where `Picocrypt.go` is located and build from source:
- Windows: <code>go build -ldflags "-s -w -H=windowsgui -extldflags=-static" Picocrypt.go</code>
- macOS: <code>go build -ldflags "-s -w -extldflags=-static" Picocrypt.go</code>
- Linux: <code>go build -ldflags "-s -w" Picocrypt.go</code>

7. You should now see a built executable (`Picocrypt.exe`/`Picocrypt.app`/`Picocrypt`) in your directory. If you're on Windows, go and download the `sdelete64.exe` from the list above and place it in the same directory as `Picocrypt.exe`.
8. You are now complete. Go ahead and run the executable you just built. Enjoy!

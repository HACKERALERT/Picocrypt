# Running From Source (Work in progress)
If you would like to run Picocrypt from source, or an executable isn't available for your platform, you've come to the right place. Running from source is very simple, and I've made it even easier with these simple instructions. Note that the instructions are generic and will work on any platform.

# 1. Prerequisites
Linux:
```bash
sudo apt-get install -y gcc make curl git tar wget xz-utils libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libgl1-mesa-dev libxxf86vm-dev libgtk-3-dev xdg-utils && sudo apt-get install -y libglx-dev || echo "" && sudo apt-get install -y xclip || sudo apt-get install xsel
```
macOS (make sure <a href="https://brew.sh/">Homebrew</a> is installed):
```bash
sudo xcode-select --install && brew install glfw3 glew
```

Windows: No prerequisites
# 2. Install the Go Programming Language
Download the corresponding installer for Go from <a href="https://golang.org/dl">here</a>, or from your package manager. Go 1.16 or higher is recommended.

# 3. Get the Source Files
Download the source file `Picocrypt.go` from above or `git clone` this repository.

# 4. Initialize Go:
Go to where `Picocrypt.go` is located and initialize Go:
```bash
go mod init Picocrypt
```

# 5. Install Picocrypt's Go Dependencies
```bash
go get -u -v github.com/pkg/browser && go get -u -v github.com/zeebo/blake3 && go get -u -v golang.org/x/crypto/sha3 && go get -u -v golang.org/x/crypto/argon2 && go get -u -v github.com/AllenDang/giu@v0.5.4 && go get -u -v github.com/OpenDiablo2/dialog && go get -u -v golang.org/x/crypto/blake2b && go get -u -v golang.org/x/crypto/blake2s && go get -u -v github.com/atotto/clipboard && go get -u -v github.com/klauspost/reedsolomon && go get -u -v golang.org/x/crypto/chacha20poly1305 && go get -u -v github.com/HACKERALERT/Picocypher/monocypher
```
Note: if macOS prompts you to install clang, do so and run the command again

# 6. Tidy the Modules:
```bash
go mod tidy
```
# 7. Build From Source
- Windows: <code>go build -ldflags "-s -w -H=windowsgui -extldflags=-static" Picocrypt.go</code>
- macOS: <code>go build -ldflags "-s -w" Picocrypt.go</code>
- Linux: <code>go build -ldflags "-s -w" Picocrypt.go</code>

# 8. You are now complete.
You should now see a built executable (`Picocrypt.exe`/`Picocrypt.app`/`Picocrypt`) in your directory. You can run it by double-clicking or executing it in your terminal. If you're on Windows, go and download the `sdelete64.exe` from the list above and place it in the same directory as `Picocrypt.exe`.

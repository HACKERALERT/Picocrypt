# Running From Source (Work in progress)
If you would like to run Picocrypt from source, or an executable isn't available for your platform, you've come to the right place. Running from source is very simple, and I've made it even easier with these simple instructions. I'll assume that you are familiar with the Go language.

Note that the instructions are generic and will work on any platform.

# 1. Prerequisites
Linux:
```bash
sudo apt-get install -y gcc make libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libgl1-mesa-dev libxxf86vm-dev libgtk-3-dev xdg-utils libglu1-mesa xclip coreutils
```
macOS:
```bash
xcode-select --install
```
Windows: [TDM-GCC](https://jmeubank.github.io/tdm-gcc/)

# 2. Install the Go Programming Language
Download the corresponding installer for Go from <a href="https://golang.org/dl">here</a>, or from your package manager. Go 1.16 or higher is recommended.

# 3. Get the Source Files
Download the source file `Picocrypt.go` from above or `git clone` this repository.

# 4. Get Resources
Place all the files in the `resources/` directory in the same directory as `Picocrypt.go`.

# 5. Create a go.mod:
Go to where `Picocrypt.go` is located and create a go.mod:
```bash
go mod init Picocrypt
```

# 5. Install Picocrypt's Go Dependencies
Install each of the following modules via `go get -u -v `:
```bash
golang.org/x/crypto
github.com/HACKERALERT/serpent@v0.0.0-20210716182301-293b29869c66
github.com/HACKERALERT/infectious@v0.0.0-20210730231340-8af02cb9ed0a
github.com/HACKERALERT/clipboard@v0.1.5-0.20210716140604-61d96bf4fc94
github.com/HACKERALERT/dialog@v0.0.0-20210716143851-223edea1d840
github.com/HACKERALERT/browser@v0.0.0-20210730230128-85901a8dd82f
github.com/HACKERALERT/zxcvbn-go@v0.0.0-20210730224720-b29e9dba62c2
```

# 6. Tidy the Modules:
```bash
go mod tidy
```

# 7. Build From Source
- Windows: <code>go build -ldflags "-s -w -H=windowsgui -extldflags=-static" Picocrypt.go</code>
- macOS: <code>go build -ldflags "-s -w" Picocrypt.go</code>
- Linux: <code>go build -ldflags "-s -w" Picocrypt.go</code>

# 8. You are now complete.
You should now see a built executable (`Picocrypt.exe`/`Picocrypt`) in your directory. You can run it by double-clicking or executing it in your terminal.

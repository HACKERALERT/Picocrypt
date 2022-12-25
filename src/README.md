# Running From Source
If you would like to run Picocrypt from source, or an executable isn't available for your platform, you've come to the right place. Running from source is very simple, and I've made it even easier with these straightforward instructions. Unlike VeraCrypt, which requires complex build procedures, SDKs, and assemblers, Picocrypt can easily be compiled from source with only a Go and C compiler. All you need is ten minutes and an Internet connection.

# 1. Prerequisites
**Linux:**
```bash
apt install -y gcc xorg-dev libgtk-3-dev libgl1-mesa-dev libglu1-mesa
```
**macOS:**
```bash
xcode-select --install
brew install glfw glew
```
**Windows:** A C compiler, ideally TDM-GCC or MinGW-w64

# 2. Install Go
If you don't have Go installed, download it from <a href="https://go.dev/dl/">here</a> or install it from your package manager (`apt install golang-go`). The latest version of Go is recommended, although you may fall back to Go 1.19 should any issues arise in the future.

# 3. Get the Source Files
Download the source files as a zip from the homepage or `git clone` this repository. Next, navigate to the `src/` directory, where you will find the source file (`Picocrypt.go`). You will need this file, along with `go.mod` and `go.sum`, to compile Picocrypt.

# 4. Build From Source
Finally, build Picocrypt from source:
- Windows: <code>go build -ldflags="-s -w -H=windowsgui -extldflags=-static" Picocrypt.go</code>
- macOS: <code>go build -ldflags="-s -w" Picocrypt.go</code>
- Linux: <code>go build -ldflags="-s -w" Picocrypt.go</code>

Note: Make sure to set `CGO_ENABLED=1` if it isn't already.

# 5. Done!
You should now see a compiled executable (`Picocrypt.exe`/`Picocrypt`) in your directory. You can run it by double-clicking or executing it in your terminal. That wasn't too hard, right? Enjoy!

Note: On Linux, if hardware OpenGL isn't available, you can set `LIBGL_ALWAYS_SOFTWARE=1` to force Mesa to use software rendering. This way, Picocrypt will be able to run regardless of driver support and can even run without a GPU at all. You may also need to set `NO_AT_BRIDGE=1` to disable the accessibility bus which is known to cause potential issues.

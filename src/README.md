# Running From Source
If you would like to run Picocrypt from source, or an executable isn't available for your platform, you've come to the right place. Running from source is very simple, and I've made it even easier with these straightforward instructions. All you need is ten minutes and an Internet connection.

# 1. Prerequisites
Linux:
```bash
apt install -y gcc make libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libgl1-mesa-dev libxxf86vm-dev libgtk-3-dev xdg-utils libglu1-mesa xclip coreutils
```
macOS:
```bash
xcode-select --install
```
Windows: A C compiler, ideally [TDM-GCC](https://jmeubank.github.io/tdm-gcc/)

# 2. Install the Go Programming Language
If you don't have Go installed, download the corresponding installer for Go from <a href="https://golang.org/dl">here</a>, or from your package manager. The latest version of Go is required.

# 3. Get the Source Files
Download the source files as a zip from the homepage or `git clone` this repository.

# 4. Build From Source
- Windows: <code>go build -ldflags "-s -w -H=windowsgui -extldflags=-static"</code>
- macOS: <code>go build -ldflags "-s -w"</code>
- Linux: <code>go build -ldflags "-s -w"</code>

# 5. Done!
You should now see a compiled executable (`Picocrypt.exe`/`Picocrypt`) in your directory. You can run it by double-clicking or executing it in your terminal. That wasn't too hard, right? Enjoy!

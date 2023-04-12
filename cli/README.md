# CLI

Before you dive in, keep in mind that the CLI is limited in functionality and not meant to replace the GUI in any remote way. It only works with volumes that don't use any keyfiles or advanced features, and you will still need the GUI to do anything more than basic file encryption. You should only use the CLI when you are not able to run the GUI or need an automatable interface for encrypting and decrypting files.

# Installation
If you don't have Go installed, download it from <a href="https://go.dev/dl/">here</a> or install it from your package manager. Then, run the command below:
```bash
go install github.com/HACKERALERT/Picocrypt/cli/picocrypt@latest
```
You should now be able to run `picocrypt` in your terminal. If not, run `export PATH=$PATH:$(go env GOPATH)/bin` and try again.

# Usage
```
picocrypt -p password <file>
```
The CLI is designed to do one thing and one thing only: encrypt and decrypt a single file. Its goal isn't to be full-blown encryption tool, but to provide the basics of file encryption so that you can do the rest. This allows you to write custom scripts to encrypt your weekly backups, secure client files on a server, etc.

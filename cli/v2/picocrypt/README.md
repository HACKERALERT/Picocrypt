# Installation
If you don't have Go installed, download it from <a href="https://go.dev/dl/">here</a> or install it from your package manager. Then, run the command below:
```
go install github.com/HACKERALERT/Picocrypt/cli/v2/picocrypt@latest
```
You should now be able to run `picocrypt` in your terminal. If not, run `export PATH=$PATH:$(go env GOPATH)/bin` and try again.
# Usage
```
C:\Users\Evan>picocrypt
Usage: picocrypt [-p]aranoid [-r]eedsolo <item1> [<item2> ...]
Items: can be a file (cat.png), folder (./src), or glob (*.txt)
```
## Examples
To encrypt a single file:
```
picocrypt secret.pdf
```
To encrypt all files in the current working directory:
```
picocrypt *
```
To encrypt all PNGs and JPGs with paranoid mode and Reed-Solomon:
```
picocrypt -p -r *.png *.jpg
```
To decrypt a volume:
```
picocrypt volume.pcv
```

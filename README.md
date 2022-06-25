## tl
A tiny translation utility.

![image](https://user-images.githubusercontent.com/71683721/175786953-eb7340a2-f62c-4c3b-af80-cefd2df67b2d.png)

## Installation
Install Go according your platform by following https://go.dev/doc/install guide.

Or manually (GNU/Linux)
```sh
# Debian (or Ubuntu, Kali etc)
# Note: On Debian you'll get bit older version of go, to get bleeding edge check out https://go.dev/doc/install
apt install golang

# Arch (Manjaro, BlackArch)
# Note: Arch and all Arch based distro will pull latest version of Go, which might be unstable
pacman -Syy golang

# FreeBSD
# Note: Might contain very old version, try to build from source by following https://freebsd.sh/go/
pkg install go
```
After installation try using `go version` on terminal
If the output like this, you're good to go!
```
go version gox.x.x [arch]
```
Here *x* is used to specify the version and *arch* is used to specify the architecture that the OS supports.

#### Install make
```sh
# Debian, Arch, FreeBSD
apt install make
pacman -Syy golang
pkg install go
```

#### Install tl
##### Prebuild binary
Head over to [releases](https://github.com/Ruzie/urban/releases) and grab the lastest build

##### Build from source
* Clone the repository
```sh
git clone https://github.com/Ruzie/tl.git
```
* Move there and run `make`, meanwhile you should have `tl`

## Usage Guide
```
./tl v0.1

Usage: 
./tl [options/none] [message]

Options:
 -h, shows this message
 -l, language code translated to
 -f, file to translate
```
### Parameters usage

* Translate to English (no params)
```sh
./tl Acesta este un text de probă
```
* Translate to a specific language
```sh
./tl -l ru Acesta este un text de probă
```
* Translate a text document to English (no params)
```sh
./tl -f sample.txt
```

* Translate a text document to a specific language
```sh
./tl -l ru -f sample.txt
```
Note: In language parameter you can't pass country name, Google translate API only supports ISO language codes.

For all supported language list, check it out https://www.labnol.org/code/19899-google-translate-languages

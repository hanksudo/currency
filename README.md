# bot-currency

Simple program to response currecny rate

## Installation

### Ubuntu

```bash
sudo apt-get install binutils bison gcc
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.5.2 --binary
gvm use 1.5.2
export GOPATH=$HOME/go
export GOROOT=`go env GOROOT`
export PATH=$PATH:$GOPATH/bin

go get -u github.com/hanksudo/bot-currency
```

# bot-currency

Currecny rate from BOT (Bank Of Taiwan)

- Slack Command
- Simple API
- CSV history

## Installation

### Ubuntu

```bash
sudo apt-get install binutils bison gcc
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.7.1 --binary
gvm use 1.7.1
export GOPATH=$HOME/go
export GOROOT=`go env GOROOT`
export PATH=$PATH:$GOPATH/bin

go get -u github.com/hanksudo/bot-currency
```

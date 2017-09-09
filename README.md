# bot-currency

Currecny rate from BOT (Bank Of Taiwan)

- Slack slash command to get currency rate 

<img src="./screenshots/slash_command.png" width="320">
    
- Web API to contrl renew time
- History CSV files

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


### Cronjob - backup CSV to dropbox

```bash
# Every three hour on weekday
0 */3 * * 1,2,3,4,5 /usr/bin/python $GOPATH/src/github.com/hanksudo/bot-currency/scripts/backup_to_dropbox.py
```

### CronJob - fetch new currency data

```bash
# Every one hour on weekday
0 */1 * * 1,2,3,4,5 /opt/bin/cronic /usr/bin/curl "http://localhost:3030/renew"
```
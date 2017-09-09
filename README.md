# bot-currency

Currecny rate from BOT (Bank Of Taiwan)


## Features

- Slack slash command to get currency rate 

<img src="./screenshots/slash_command.png" width="320">
    
- Web API to contrl renew time
- History CSV files

## Installation

```bash
go get -u github.com/hanksudo/bot-currency
```

* I develop on Mac and deploy this process on Ubuntu 12.04, you can use [rvm](https://github.com/moovweb/gvm) easily manage Go versions.


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
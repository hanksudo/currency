# bot-currency

Currecny rate from BOT (Bank Of Taiwan)

## Features

- Download currency rete file from BOT
- Periodically download new currency rate CSV file.
- Periodically backup to dropbox.
- Use [Slack Slash commands](https://api.slack.com/interactivity/slash-commands) to get latest currency rate

<img src="./screenshots/slash_command.png" width="340">

## Installation

```bash
go get -u github.com/hanksudo/bot-currency
```

## Usage

```bash
# start web server for Slack
bot-currency -web

# Retrieved latest CSV file immediately
bot-currency -renew

# Backup to dropbox
export DROPBOX_ACCESS_TOKEN=<YOUR-ACCESS-TOKEN>
bot-currency -backup
```

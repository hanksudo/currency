# Currency

Currency services

## Sources

- BOT (Bank Of Taiwan) - <https://rate.bot.com.tw/xrt>

## Features

- Download currency rate CSV file from BOT
- Periodically download new currency rate CSV file.
- Periodically backup to dropbox.
- Use [Slack Slash commands](https://api.slack.com/interactivity/slash-commands) to get latest currency rate

<img src="./screenshots/slash_command.png" width="340">

## Installation

```bash
go install github.com/hanksudo/currency@latest
```

## Usage

```bash
# start web server for Slack
currency -web
curl "http://localhost:3030?currency=jpy"

# Retrieved latest CSV file
currency -renew

# Backup to Dropbox
# You need generate access token from your Dropbox app setting page
export DROPBOX_ACCESS_TOKEN=<YOUR-ACCESS-TOKEN>
currency -backup
```

## Test slack command on local environment

```bash
go run . -web
ngrok http 3030
```

Then set your ngrok url on you slack command integration url.

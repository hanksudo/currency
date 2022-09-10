# bot-currency

Currecny rate from BOT (Bank Of Taiwan)
<https://rate.bot.com.tw/xrt>

## Features

- Download currency rate CSV file from BOT
- Periodically download new currency rate CSV file.
- Periodically backup to dropbox.
- Use [Slack Slash commands](https://api.slack.com/interactivity/slash-commands) to get latest currency rate

<img src="./screenshots/slash_command.png" width="340">

## Installation

```bash
go install github.com/hanksudo/bot-currency@latest
```

## Usage

```bash
# start web server for Slack
bot-currency -web
curl "http://localhost:3030?currency=jpy"

# Retrieved latest CSV file
bot-currency -renew

# Backup to dropbox
# You need generate access token from your Dropbox app setting page
pip install -r scripts/requirements.txt
export DROPBOX_ACCESS_TOKEN=<YOUR-ACCESS-TOKEN>
bot-currency -backup
```

## Test slack command on local environment

```bash
go run . -web
ngrok http 3030
```

Then set your ngrok url on you slack command integration url.

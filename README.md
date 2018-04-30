# bot-currency

Currecny rate from BOT (Bank Of Taiwan)

## Features

- Use Slack Slash commands to get currency rate

<img src="./screenshots/slash_command.png" width="340">

- History CSV files

## Installation

```bash
go get -u github.com/hanksudo/bot-currency
```

- I develop on Mac and deploy this process on Ubuntu 12.04, you can use [gvm](https://github.com/moovweb/gvm) easily manage Go versions.

## Usage

```bash
export DROPBOX_ACCESS_TOKEN=<YOUR-ACCESS-TOKEN>
bot-currency -web
bot-currency -renew
bot-currency -backup
```
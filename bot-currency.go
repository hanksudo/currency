package main

import (
	"flag"
	"os"

	"github.com/hanksudo/bot-currency/backup"
	"github.com/hanksudo/bot-currency/currency"
	"github.com/hanksudo/bot-currency/web"
)

// RootPath - root of project
var RootPath string

func main() {
	RootPath, _ = os.Getwd()
	webPtr := flag.Bool("web", false, "Start web server")
	renewPtr := flag.Bool("renew", false, "Renew currency data")
	backupPtr := flag.Bool("backup", false, "Backup to Dropbox")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
	} else if *webPtr {
		web.Start()
	} else if *renewPtr {
		currency.Renew()
	} else if *backupPtr {
		backup.Start()
	}
}

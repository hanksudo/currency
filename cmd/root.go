package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hanksudo/currency/backup"
	"github.com/hanksudo/currency/config"
	"github.com/hanksudo/currency/handlers"
	"github.com/hanksudo/currency/services/bot"
	"github.com/robfig/cron/v3"
)

func Start() {
	versionPtr := flag.Bool("version", false, "Print the version")
	webPtr := flag.Bool("web", false, "Start web server")
	renewPtr := flag.Bool("renew", false, "Renew currency data")
	backupPtr := flag.Bool("backup", false, "Backup to Dropbox")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
	} else if *versionPtr {
		fmt.Printf("%s", config.VERSION)
	} else if *webPtr {
		c := cron.New()
		// Renew - Every one hour on weekday
		c.AddFunc("0 0 * * * 1,2,3,4,5", bot.Renew)
		// Backup - Every 3 hours on weekday
		c.AddFunc("0 0 */3 * * 1,2,3,4,5", backup.Start)
		c.Start()

		handlers.Start()
	} else if *renewPtr {
		bot.Renew()
	} else if *backupPtr {
		backup.Start()
	} else {
		log.Println("No command match!")
	}
}

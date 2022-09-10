package backup

import (
	"fmt"
	"log"
	"os"

	"github.com/hanksudo/bot-currency/services"
	"golang.org/x/exp/slices"
)

// Start - backup
func Start() {
	log.Println("Start backup files to Dropbox...")
	token := os.Getenv("DROPBOX_ACCESS_TOKEN")
	if token == "" {
		log.Fatalln("Stop backup, DROPBOX_ACCESS_TOKEN missing.")
		os.Exit(2)
	}

	service := services.NewDropboxService(token)
	if !service.Authenticated() {
		log.Fatalln("Dropbox unauthenticated!")
	}

	existsFilenames := service.ListAllFilenames()

	rootPath, _ := os.Getwd()
	folderPath := fmt.Sprintf("%s/csvs", rootPath)
	files, err := os.ReadDir(folderPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !slices.Contains(existsFilenames, file.Name()) {
			f, err := os.Open(fmt.Sprintf("%s/%s", folderPath, file.Name()))
			if err != nil {
				panic(err)
			}
			log.Println("Uploading file:", file.Name())
			service.UploadFile(f, file.Name())
		}
	}
}

package services

import (
	"fmt"
	"log"
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/users"
)

type DropboxService struct {
	Config dropbox.Config
}

func NewDropboxService(token string) *DropboxService {
	return &DropboxService{
		Config: dropbox.Config{
			Token:    token,
			LogLevel: dropbox.LogInfo, // if needed, set the desired logging level. Default is off
		},
	}
}

func (s *DropboxService) Authenticated() bool {
	_, err := s.NewUserClient().GetCurrentAccount()
	return err == nil
}

func (s *DropboxService) UploadFile(file *os.File, filename string) {
	dbx := s.NewFileClient()

	_, err := dbx.Upload(files.NewUploadArg(fmt.Sprintf("/%s", filename)), file)
	if err != nil {
		log.Fatal(err)
	}
}
func (s *DropboxService) ListAllFilenames() []string {
	dbx := s.NewFileClient()

	res, err := dbx.ListFolder(&files.ListFolderArg{})
	if err != nil {
		panic(err)
	}

	entries := res.Entries
	for res.HasMore {
		arg := files.NewListFolderContinueArg(res.Cursor)
		res, err = dbx.ListFolderContinue(arg)
		if err != nil {
			panic(err)
		}
		entries = append(entries, res.Entries...)
		if !res.HasMore {
			break
		}
	}

	var filenames []string
	for _, entry := range entries {
		switch entry := entry.(type) {
		case *files.FileMetadata:
			filenames = append(filenames, entry.Name)
		default:
			log.Println("This entry not FileMetadata", entry)
		}

	}
	return filenames
}

func (s *DropboxService) NewUserClient() users.Client {
	return users.New(s.Config)
}

func (s *DropboxService) NewFileClient() files.Client {
	return files.New(s.Config)
}

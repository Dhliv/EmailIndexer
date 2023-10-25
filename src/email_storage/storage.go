package email_storage

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/dhliv/EmailIndexing/src/commands"
	"github.com/dhliv/EmailIndexing/src/constants"
	"github.com/dhliv/EmailIndexing/src/parser"
	"github.com/dhliv/EmailIndexing/src/requests"
)

type emailStorage struct {
	bulkSize int
	emailsFile *os.File
	lock *sync.Mutex
}

var storage *emailStorage = nil

func init() {
	emailsFile, err := createNDJSONFile()
	if err != nil {
		log.Printf("Error while creating ndjson file: %v\n", err)
		return
	}

	storage = &emailStorage{
		bulkSize: 0,
		lock: &sync.Mutex{},
		emailsFile: emailsFile,
	}
}

// Sync Safe function. Adds email to storage, uploading all emails in storage
// in case the size does exceed the maximum established.
func AddEmailToStorage(email map[string]*string) {
	storage.lock.Lock()
	defer storage.lock.Unlock()
	size, err := strconv.Atoi(*email["size"])
	if err != nil {
		log.Printf("Error while casting string to int for email size: %v\n", err)
		return
	}
	delete(email, "size")

	if size > constants.MAX_EMAIL_SIZE {
		err := requests.UploadSingleEmailRecord(email)
		if err != nil {
			log.Printf("Error while uploading single email: %v\n", err)
		}

		return
	}

	hasExceededBulkSize := size + storage.bulkSize > constants.MAX_STORAGE_SIZE
	if hasExceededBulkSize{
		err := UploadBulkEmails()
		if err != nil {
			log.Printf("Error while bulk uploading emails to ZincSeacrh: %v\n", err)
			return
		}
	}

	ndjsonString, err := parser.EmailMapToNDJSON(email)
	if err != nil {
		log.Printf("Error encountered while marshaling map: %v\n", email)
		return
	}
	
	_, err = storage.emailsFile.Write([]byte(*ndjsonString))
	if err != nil {
		log.Printf("Error while writing email to emailFile: %v\n", err)
		return
	}

	storage.bulkSize += size
}

/*
Bulk Uploads all emails in 'storage.emailsFile' and resets storage.
*/
func UploadBulkEmails() error {
	err := commands.BulkUploadToZincSearch(storage.emailsFile.Name())
	if err != nil {
		return err
	}
	
	resetStorage()
	return nil
}

func createNDJSONFile() (*os.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	emailsFilePath := filepath.Join(wd, "emails.ndjson")
	emailsFile, err := os.Create(emailsFilePath)
	if err != nil {
		return nil, err
	}

	return emailsFile, nil
}

/*
Deletes 'storage.emailsFile', creates a new one and
resets storage with a new empty emails file and bulkSize=0.
*/
func resetStorage() {
	err := os.Remove(storage.emailsFile.Name())
	if err != nil {
		log.Printf("Error while removing ndjson file: %v\n", err)
		return
	}

	file, err := createNDJSONFile()
	if err != nil {
		log.Printf("Error while creating ndjson file: %v\n", err)
		return
	}

	storage.bulkSize = 0
	storage.emailsFile = file
}
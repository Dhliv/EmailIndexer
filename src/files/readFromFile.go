package files

import (
	"fmt"
	"os"

	"github.com/dhliv/EmailIndexing/src/constants"
	"github.com/dhliv/EmailIndexing/src/email_storage"
	"github.com/dhliv/EmailIndexing/src/parser"
)

/*
Reads text from file in ´filePath' and parse it to email map and passes it
to email storage.
*/
func readEmailFromFile(filePath string) error{
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	if fileInfo.Size() > constants.MAX_FILE_SIZE {
		return fmt.Errorf("File %v exceeds maximum file size %v\n", filePath, constants.MAX_FILE_SIZE)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	emailParser := parser.NewEmailParser(file)
	res := emailParser.ParseEmailFile()
	sizeString := fmt.Sprint(fileInfo.Size())
	res["size"] = &sizeString

	go email_storage.AddEmailToStorage(res)

	return nil
}
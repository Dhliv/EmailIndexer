package commands

import (
	"fmt"
	"os/exec"

	"github.com/dhliv/EmailIndexing/src/constants"
)

func BulkUploadToZincSearch(filePath string) error {
	command := exec.Command("curl", fmt.Sprintf("%v/api/_bulk", constants.API_URL), "-i", "-u",
													fmt.Sprintf("%v:%v", constants.USER_NAME, constants.USER_PASSWORD),
													"--data-binary", fmt.Sprintf("@%v", filePath))
	err := command.Run()
	if err != nil {
		return err
	}

	return nil
}
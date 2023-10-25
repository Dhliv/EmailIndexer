package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/dhliv/EmailIndexing/src/constants"
)

func UploadSingleEmailRecord(email map[string]*string) error {
	emailJson, err := json.Marshal(email)
	if err != nil {
		return err
	}
	
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/api/%v/_doc", constants.API_URL, constants.API_INDEX), strings.NewReader(string(emailJson)))
	if err != nil {
		return err
	}
	req.SetBasicAuth(constants.USER_NAME, constants.USER_PASSWORD)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
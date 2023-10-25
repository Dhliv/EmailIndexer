package parser

import (
	"encoding/json"
	"fmt"

	"github.com/dhliv/EmailIndexing/src/constants"
)

func EmailMapToNDJSON(email map[string]*string) (*string, error) {
	resString := fmt.Sprintf("{ \"index\" : { \"_index\" : \"%v\" } }\n", constants.API_INDEX)
	res, err := json.Marshal(email)
	if err != nil {
		return nil, err
	}

	resString += string(res) + "\n"

	return &resString, nil
}
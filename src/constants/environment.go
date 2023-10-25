package constants

import (
	"log"
	"os"
)

var API_INDEX string = "API_INDEX"
var USER_NAME string = "USER_NAME"
var USER_PASSWORD string = "USER_PASSWORD"
var API_URL string = "API_URL"

func getVar(varName string) string {
	newVal, hasVar := os.LookupEnv(varName)
	if !hasVar {
		log.Printf("Environment variable %v it's not set!\n", varName)
	}

	return newVal
}

func init() {
	API_INDEX = getVar(API_INDEX)
	USER_NAME = getVar(USER_NAME)
	USER_PASSWORD = getVar(USER_PASSWORD)
	API_URL = getVar(API_URL)
}
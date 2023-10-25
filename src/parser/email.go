package parser

import (
	"bufio"
	"os"
	"strings"
)

var emailFields []string = []string{"Message-ID:", "Date:", "From:", "To:", "Subject:", "Mime-Version:",
																		"Content-Type:", "Content-Transfer-Encoding:", "X-From:", "X-To:",
																		"X-cc:", "X-bcc:", "X-Folder:", "X-Origin:", "X-FileName:", "Body"}

type EmailParser struct {
	content string
	fieldsObtained map[string]*string
	emailFieldPosition int
	fileScanner *bufio.Scanner
}


// EmailParser constructor. Creates a parser assigned to file.
func NewEmailParser(file *os.File) *EmailParser {
	return &EmailParser{
		content: "",
		fieldsObtained: make(map[string]*string),
		emailFieldPosition: 0,
		fileScanner: bufio.NewScanner(file),
	}
}


// Parses file into email map
func (ep *EmailParser) ParseEmailFile() map[string]*string {
	ep.parseHeaders()
	ep.parseBody()
	return ep.fieldsObtained
}


// Reads all explicit fields in email file, except from Body.
func (ep *EmailParser) parseHeaders() {
	ep.parseFirstField()
	for ep.fileScanner.Scan() {
		s := ep.fileScanner.Text()

		if len(s) == 0 { // has only endline
			break
		}

		cuttedS, hasEmailField := strings.CutPrefix(s, emailFields[ep.emailFieldPosition])
		if !hasEmailField {
			if ep.content != "" {
				ep.content += "\n"
			}

			ep.content += cuttedS
			continue
		}

		ep.addEmailField()
		ep.content = cuttedS
	}
	
	ep.addEmailField()
	ep.content = ""
}


func (ep *EmailParser) parseBody() {
	for ep.fileScanner.Scan() {
		s := ep.fileScanner.Text()
		if ep.content != "" {
			ep.content += "\n"
		}

		ep.content += strings.Trim(s, " ")
	}

	ep.addEmailField()
}


func (ep *EmailParser) addEmailField() {
	emailField, _ := strings.CutSuffix(emailFields[ep.emailFieldPosition - 1], ":")
	s := strings.Trim(ep.content, " ")
	ep.fieldsObtained[emailField] = &s
	ep.emailFieldPosition++
}


func (ep *EmailParser) parseFirstField() {
	ep.fileScanner.Scan()
	s := ep.fileScanner.Text()
	ep.content, _ = strings.CutPrefix(s, emailFields[ep.emailFieldPosition])
	ep.emailFieldPosition++
}
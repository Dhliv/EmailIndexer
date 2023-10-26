package parser

import (
	"bufio"
	"os"

	"github.com/dhliv/EmailIndexing/src/fast_strings"
)

var emailFields []string = []string{"Message-ID:", "Date:", "From:", "To:", "Subject:", "Mime-Version:",
																		"Content-Type:", "Content-Transfer-Encoding:", "X-From:", "X-To:",
																		"X-cc:", "X-bcc:", "X-Folder:", "X-Origin:", "X-FileName:", "Body"}
var emailFieldsCurated []string = []string{"Message-ID", "Date", "From", "To", "Subject", "Mime-Version",
																						"Content-Type", "Content-Transfer-Encoding", "X-From", "X-To",
																						"X-cc", "X-bcc", "X-Folder", "X-Origin", "X-FileName", "Body"}

var endline string = "\n"

type EmailParser struct {
	content *fast_strings.FastString
	fieldsObtained map[string]*string
	emailFieldPosition int
	fileScanner *bufio.Scanner
}


// EmailParser constructor. Creates a parser assigned to file.
func NewEmailParser(file *os.File) *EmailParser {
	return &EmailParser{
		content: fast_strings.NewFastString(),
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
		fs := fast_strings.NewFastString()

		if len(s) == 0 { // has only endline
			break
		}

		fs.Concat(&s)
		hasEmailField := fs.CutPrefix(&emailFields[ep.emailFieldPosition])
		if !hasEmailField {
			if ep.content.Size() != 0 {
				ep.content.Concat(&endline)
			}

			ep.content.ConcatFastString(fs)
			continue
		}

		ep.addEmailField()
		ep.content = fs
	}
	
	ep.addEmailField()
	ep.content = fast_strings.NewFastString()
}


func (ep *EmailParser) parseBody() {
	for ep.fileScanner.Scan() {
		s := ep.fileScanner.Text()
		if ep.content.Size() != 0 {
			ep.content.Concat(&endline)
		}

		ep.content.Concat(&s)
	}

	ep.addEmailField()
}


func (ep *EmailParser) addEmailField() {
	emailField := &emailFieldsCurated[ep.emailFieldPosition - 1]
	ep.fieldsObtained[*(emailField)] = ep.content.GetString()
	ep.emailFieldPosition++
}


func (ep *EmailParser) parseFirstField() {
	ep.fileScanner.Scan()
	s := ep.fileScanner.Text()
	fs := fast_strings.NewFastString()
	fs.Concat(&s)
	fs.CutPrefix(&emailFields[ep.emailFieldPosition])
	ep.content = fs
	ep.emailFieldPosition++
}
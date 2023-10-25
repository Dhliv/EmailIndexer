package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"

	"github.com/dhliv/EmailIndexing/src/email_storage"
	"github.com/dhliv/EmailIndexing/src/files"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		fmt.Println("Profiling")
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	args := os.Args

	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Indexer should include the folder to index: ./index.sh enron_mail_20110402")
		return
	}

	path, err := os.Getwd()
	if err != nil {
		log.Printf("Error while getting working directory: %v\n", err)
		return
	}

	directory := args[len(args) - 1]
	directoryPath := filepath.Join(path, directory)
	err = files.WalkThroughFolders(directoryPath)
	if(err != nil) {
		log.Printf("Error while walking through directory %v: %v\n", path, err)
		return
	}

	email_storage.UploadBulkEmails()
}
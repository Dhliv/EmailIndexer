package files

import (
	"io/fs"
	"log"
	"path/filepath"
)

/*
Reads all files (not folders) into emails.
*/
func filterFolders(path string, info fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	err = readEmailFromFile(path)
	if err != nil {
		log.Printf("Error while reading from file %v: %v\n", path, err)
		return nil
	}

	return nil
}

/*
	Walks through every folder recursively and applys 'filterFolders' function
	in every file and folder found
*/
func WalkThroughFolders(directory string) error {
	err := filepath.WalkDir(directory, filterFolders)
	if err != nil {
		return err
	}

	return nil
}
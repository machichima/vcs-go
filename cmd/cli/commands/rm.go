package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

func executeRm(filePath string) error {

	var isRm bool = false

	// standarize the format of the filePath
	filePath = filepath.Join(filePath)

	index, err := utils.ReadIndexFile()
	if err != nil {
        if !os.IsNotExist(err) {
            return err
        }
	} else {

		// check if the filePath is in index
		if _, ok := index.FileToHash[filePath]; ok {
			// filePath in index
			delete(index.FileToHash, filePath)
			isRm = true
		}

		// check if the filePath is the dir of the files
		// in the index
		for file, _ := range index.FileToHash {
			matched, err := filepath.Match(filepath.Join(filePath, "*"), file)
			if err != nil {
				return err
			}

			if matched {
				delete(index.FileToHash, file)
				isRm = true
			}
		}

		if err := utils.WriteIndexFile(index); err != nil {
			return err
		}

	}

	if isRm {
		fmt.Println("Files unstaged")
	} else {
		fmt.Println("No files unstaged")
	}

	return nil
}

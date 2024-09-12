package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

// steps:
// 1. check if index file exists (if not, print no staged and return)
// 2. get the bytes content in the index file
// 3. deserialize the bytes content into Index struct
func executeStatus() error {

    index, err := utils.ReadIndexFile()
    if err != nil {
        if !errors.Is(err, os.ErrNotExist) {
            return err
        }
    }

    fmt.Println("Changes to be committed:")
    for file, _ := range index.FileToHash {
        fmt.Println(file)
    }

    files, err := utils.GetFiles("./")
    if err != nil {
        return err
    }

    fmt.Println("Changes not staged for commit:")
    for _, file := range files {
        _, ok := index.FileToHash[file]
        if !ok {
            fmt.Println(file)
        }
    }

    return nil
}

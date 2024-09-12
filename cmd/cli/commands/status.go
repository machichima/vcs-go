package commands

import (
	"bytes"
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

    indexBytes, err := os.ReadFile(utils.IndexDirName)
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("No staged changes")
        return nil
    }

    indexBuf := bytes.NewBuffer(indexBytes)
    index, err := utils.DeserializeIndex(indexBuf)
    if err != nil {
        return err
    }

    if len(index.FileToHash) < 1 {
        fmt.Println("No staged changes")
        return nil
    }

    fmt.Println("Staged files:")
    for file, _ := range index.FileToHash {
        fmt.Println(file)
    }


    return nil
}

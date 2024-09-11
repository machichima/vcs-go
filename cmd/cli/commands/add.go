package commands

import (
	// "errors"
	"bytes"
    "fmt"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

// steps:
//  1. go through all files, get list of files
//  2. serialize the files
//  3. hash the files to get sha-1 hash
//  4. put files to objects dir (based on hash)
func executeAdd(filePath string) error {

	// TODO: distinguish between file and dir (now dir only)

	files, err := utils.GetFiles(filePath)
	if err != nil {
		return err
	}

    var isAddNewFile bool = false

	for _, file := range files {
		blobStruct, err := utils.FileToStruct(file)
		if err != nil {
			return err
		}

		// Serialize the blob
		var buffer bytes.Buffer
		if err := utils.SerializeBlob(blobStruct, &buffer); err != nil {
			return err
		}

		// hash the serialized blob
		hashStr, err := utils.HashBlob(buffer.Bytes())
		if err != nil {
			return err
		}

        isNewFile, err := utils.SaveFileByHash(file, hashStr, buffer.Bytes(), utils.AddType)
        if err != nil {
            return err
        }

        isAddNewFile = isAddNewFile || isNewFile
	}

    if isAddNewFile {
        fmt.Println("Files added successfully")
    } else {
        fmt.Println("No new files added")
    }

	return nil
}

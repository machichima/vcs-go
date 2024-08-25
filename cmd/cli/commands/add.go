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

	var hashBlob = make(map[string][]byte)
	var fileHash = make(map[string]string)

	for _, file := range files {
		blobStruct, err := utils.FileToStruct(file)
		if err != nil {
			return err
		}

		// Serialize the blob
		var buffer bytes.Buffer
		if err := utils.Serialize(blobStruct, &buffer); err != nil {
			return err
		}

		// hash the serialized blob
		hashStr, err := utils.HashBlob(buffer.Bytes())
		if err != nil {
			return err
		}

		// save hash string with blob
		hashBlob[hashStr] = buffer.Bytes()
		fileHash[file] = hashStr

		if err := utils.SaveFileByHash(hashStr, blobStruct.Bytes, utils.AddType); err != nil {
			return err
		}
	}

	fmt.Println(hashBlob)
	fmt.Println(fileHash)

	// TODO: save serialized blob to the files using following function

	// TODO: save to .vgo/object dir, and in .vgo/index file write
	// hash and file name (like the file tree)

	return nil
}

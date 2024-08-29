package utils_test

import (
	"log"
	"os"
	"path/filepath"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

const TestStr = `This is a test
second line`

func CreateTempFile(dir string) (string, error) {
	path := dir + "/test.txt"

	err := os.WriteFile(path, []byte(TestStr), os.ModePerm)
	if err != nil {
		log.Fatal("Error occur while writing temp content to file")
	}

	return path, nil
}

// Create a temp file in the .vgo/objects/hash[:2] dir
// return the full path for hash file
func CreateTempFileInVgo(dir string) (string, error) {

	hash := "dfq8hffjalkgihboq"
	filePath := filepath.Join(dir, utils.ObjectsDirName, hash[:2], hash)

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return "", err
	}

	if err := os.WriteFile(filePath, []byte(TestStr), os.ModePerm); err != nil {
		return "", err
	}

	return filePath, nil
}

// create a temp .vgo in the dir with numOfFile files
// return the repo path
func CreateTempVgo(dir string, numOfFile int) {
	// TODO: Create a temp repo
}

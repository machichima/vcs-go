package utils_test

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

func TestGetFiles(t *testing.T) {
	dirs, err := utils.GetFiles("../..")
	if err != nil {
		t.Error("Error occur while getting directories")
	}

	if len(dirs) == 0 {
		t.Error("No directories found")
	}
}

func TestFileToStruct(t *testing.T) {

	blob, err := utils.FileToStruct("./test.txt")
	if err != nil {
		t.Error("Error occur while converting file to struct")
	}

	if reflect.TypeOf(blob) != reflect.TypeOf(utils.Blob{}) {
		t.Errorf("type %v %v mismatch", reflect.TypeOf(blob), reflect.TypeOf(utils.Blob{}))
	}

	if blob.Bytes == nil {
		t.Error("Bytes is nil")
	}
}

func TestSaveFileByHash(t *testing.T) {
    dir := t.TempDir()

    path, err := CreateTempFile(dir)
    if err != nil {
        t.Error("Error occur while creating temp file")
        t.Error(err)
    }

	fileByte, err := os.ReadFile(path)
	if err != nil {
		t.Error("Error occur while reading temp file")
	}

	hash, err := utils.HashBlob(fileByte)
	if err != nil {
		t.Error("Error occur while hashing file")
	}

    // change the working dir
    if err := os.Chdir(dir); err != nil {
        t.Error("Error occur while changing the working directory")
        t.Error(err)
    }

    if err := utils.SaveFileByHash(hash, fileByte, utils.CommitType); err != nil {
        t.Error("Error occur while saving file by hash")
        t.Error(err)
    }

    // check if the file is saved
	parentDir := hash[:2]
    fullObjectsPath := utils.ObjectsDirName + "/" + parentDir + "/" + hash

    if _, err := os.Stat(fullObjectsPath); errors.Is(err, os.ErrNotExist) {
        t.Error("File not saved")
    } else if err != nil {
        t.Error("Error occur while checking the saved file")
        t.Error(err)
    }

    content, err := os.ReadFile(fullObjectsPath)
    if string(content) != TestStr {
        t.Error("Saved file content mismatch")
    }

}

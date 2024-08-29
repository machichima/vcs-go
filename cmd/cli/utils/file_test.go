package utils_test

import (
	"errors"
	"os"
	"path/filepath"
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

func TestDeleteObject(t *testing.T) {
	dir := t.TempDir()
	if err := os.Chdir(dir); err != nil {
		t.Error(err)
	}

	path, err := CreateTempFileInVgo(dir)
	if err != nil {
		t.Error(err)
	}
    hash := filepath.Base(path)

	// ensure the object file is created
	objDir := filepath.Join(dir, utils.ObjectsDirName)
	folders, err := os.ReadDir(filepath.Join(objDir, hash[:2]))
	if err != nil {
		t.Error(err)
	}

	if len(folders) == 0 {
		t.Error("Object folder is not created")
	}

	if err := utils.DeleteObject(hash); err != nil {
		t.Error(err)
	}

	// ensure the object folder is empty
	folders, err = os.ReadDir(objDir)
	if err != nil {
		t.Error(err)
	}

	if len(folders) != 0 {
		t.Error("Object folder is not empty")
	}

}

func TestAddToIndex(t *testing.T) {

	dir := t.TempDir()
	if err := os.Chdir(dir); err != nil {
		t.Error(err)
	}

	path, err := CreateTempFileInVgo(dir)
	if err != nil {
		t.Error(err)
	}
    hash := filepath.Base(path)

	// Case 1: object not yet in INDEX
	var index utils.Index
    index.FileToHash = make(map[string]string)

    if err := utils.AddToIndex(&index, "file", hash); err != nil {
        t.Error(err)
    }

    if len(index.FileToHash) != 1 {
        t.Error("Error adding hash-file pair to index")
    }

    // Case 2: object already in INDEX - same hash
    if err := utils.AddToIndex(&index, "file", hash); err != nil {
        t.Error(err)
    }

    if len(index.FileToHash) != 1 {
        t.Error("Duplicate identical file-hash pair added")
    }

    // Case 2: object already in INDEX - different hash
    newHash := "1tjqwfwajq3j0jg3"
    if err := utils.AddToIndex(&index, "file", newHash); err != nil {
        t.Error(err)
    }

    // make sure the old object is deleted
    folders, err := os.ReadDir(filepath.Join(dir, utils.ObjectsDirName))
	if err != nil {
		t.Error(err)
	}

	if len(folders) != 0 {
		t.Error("Object folder is not empty, the old hash objects are not deleted")
	}

    if index.FileToHash["file"] != newHash {
        t.Error("File hash is not updated")
    }

}

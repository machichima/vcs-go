package utils_test

import (
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
}

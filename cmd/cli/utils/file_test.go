package utils

import (
	"reflect"
	"testing"
)

func TestGetDirs(t *testing.T) {
    dirs, err := GetDirs("../..")
    if err != nil {
        t.Error("Error occur while getting directories")
    }

    if len(dirs) == 0 {
        t.Error("No directories found")
    }
}


func TestFileToStruct(t *testing.T) {

    blob, err := FileToStruct("./test.txt")
    if err != nil {
        t.Error("Error occur while converting file to struct")
    }

    if reflect.TypeOf(blob) != reflect.TypeOf(Blob{}) {
        t.Errorf("type %v %v mismatch", reflect.TypeOf(blob), reflect.TypeOf(Blob{}))
    }
}

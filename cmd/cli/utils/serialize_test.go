package utils

import (
	"reflect"
	"testing"
)

func TestSerialize(t *testing.T) {

	// Read the file content and convert to struct
    blob, err := FileToStruct("./test.txt")
    if err != nil {
        t.Error("Error occur while converting file to struct")
    }

	serializedBytes, err := Serialize(blob)
	if err != nil {
		t.Error("Error serializing data")
	}

    if reflect.TypeOf(serializedBytes) != reflect.TypeOf([]byte{}) {
        t.Error("Error serializing data")
        t.Errorf("type %v %v mismatch", reflect.TypeOf(serializedBytes), reflect.TypeOf([]byte{}))
    }
}

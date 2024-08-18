package utils

import (
    "fmt"
	"bytes"
	"testing"
    "reflect"
)

func TestHashBlob(t *testing.T) {

	// Read the file content and convert to struct
    blob, err := FileToStruct("./test.txt")
    if err != nil {
        t.Error("Error occur while converting file to struct")
    }

    // Serialize the blob
	var buffer bytes.Buffer

	serializeErr := Serialize(blob, &buffer)
	if serializeErr != nil {
		t.Errorf("Error serializing data with error %v", serializeErr)
	}

    // hash the serialized blob
    hashBytes, err := HashBlob(buffer.Bytes())
    if err != nil {
        t.Error("Error occur while hashing blob")
    }

    fmt.Printf("Hashed blob: %v\n", hashBytes)

    if reflect.TypeOf(hashBytes).Kind() != reflect.String {
        t.Error("Error deserializing data")
        t.Errorf("type %v %v mismatch", reflect.TypeOf(hashBytes), reflect.String)
    }
}

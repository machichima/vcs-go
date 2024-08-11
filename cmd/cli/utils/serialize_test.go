package utils

import (
    "fmt"
	"bytes"
	"reflect"
	"testing"
)

func TestSerialize(t *testing.T) {

	// Read the file content and convert to struct
    blob, err := FileToStruct("./test.txt")
    if err != nil {
        t.Error("Error occur while converting file to struct")
    }

	var buffer bytes.Buffer

	serializeErr := Serialize(blob, &buffer)
	if serializeErr != nil {
		t.Errorf("Error serializing data with error %v", serializeErr)
	}

    // if reflect.TypeOf(serializedBytes) != reflect.TypeOf([]byte{}) {
    //     t.Error("Error serializing data")
    //     t.Errorf("type %v %v mismatch", reflect.TypeOf(serializedBytes), reflect.TypeOf([]byte{}))
    // }
}


func TestDeserialize(t *testing.T) {

	// Do serialize step first
    blob, err := FileToStruct("./test.txt")
    if err != nil {
        t.Error("Error occur while converting file to struct")
    }

	var buffer bytes.Buffer
    fmt.Printf("Buffer addr in TestDeserialize: %p\n", &buffer)

	serializeErr := Serialize(blob, &buffer)
	if serializeErr != nil {
		t.Errorf("Error serializing data with error %v", serializeErr)
	}

	decodedBytes, err := Deserialize(&buffer)
	if err != nil {
		t.Errorf("Error deserializing data with error %v", err)
	}

    t.Logf("Decoded bytes string %v", string(decodedBytes))

    if reflect.TypeOf(decodedBytes) != reflect.TypeOf([]byte{}) {
        t.Error("Error deserializing data")
        t.Errorf("type %v %v mismatch", reflect.TypeOf(decodedBytes), reflect.TypeOf([]byte{}))
    }
}

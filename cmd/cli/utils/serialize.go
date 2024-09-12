package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type TestData struct {
	Content string
}

// SerializeBlob the data and put back to the buffer
func SerializeBlob(data Blob, buffer *bytes.Buffer) error {
	// fmt.Printf("Buffer addr in Serialize: %p\n", buffer)

	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)

	return err
}

func SerializeIndex(data Index, buffer *bytes.Buffer) error {

	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)

	return err
}

func SerializeCommit(data Commit, buffer *bytes.Buffer) error {

	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)

	return err
}

func DeserializeBlob(buffer *bytes.Buffer) ([]byte, error) {

	var b Blob

	decoder := gob.NewDecoder(buffer)

	for {
		err := decoder.Decode(&b)
		if err != nil {
			fmt.Println("Error in decoder", err)
			break
		}
	}

	return b.Bytes, nil
}

func DeserializeIndex(buffer *bytes.Buffer) (Index, error) {

	var b Index

	decoder := gob.NewDecoder(buffer)

	for {
		err := decoder.Decode(&b)
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				fmt.Println("Error in index decoder", err)
			}
		}
	}

	return b, nil
}

func DeserializeCommit(buffer *bytes.Buffer) (Commit, error) {

	var b Commit

	decoder := gob.NewDecoder(buffer)

	for {
		err := decoder.Decode(&b)
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				fmt.Println("Error in index decoder", err)
			}
		}
	}

	return b, nil
}

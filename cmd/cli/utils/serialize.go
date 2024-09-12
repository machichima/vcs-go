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

	// fmt.Printf("Buffer addr in Deserialize: %p\n", buffer)

	// var b bytes.Buffer
	var b Blob

	decoder := gob.NewDecoder(buffer)

	// fmt.Println("Bytes to be serialized: ")
	// fmt.Println(*buffer)

	for {
		err := decoder.Decode(&b)
		if err != nil {
			fmt.Println("Error in decoder", err)
			break
		}
	}

	// fmt.Println("Decoded bytes", string(b.Bytes))

	// decodedData := b.Bytes()
	return b.Bytes, nil
}

func DeserializeIndex(buffer *bytes.Buffer) (Index, error) {

	// fmt.Printf("Buffer addr in Deserialize: %p\n", buffer)

	// var b bytes.Buffer
	var b Index

	decoder := gob.NewDecoder(buffer)

	// fmt.Println("Bytes to be serialized: ")
	// fmt.Println(*buffer)

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

	// fmt.Println("Decoded bytes", string(b.Bytes))

	// decodedData := b.Bytes()
	return b, nil
}

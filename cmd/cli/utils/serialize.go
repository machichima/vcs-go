package utils

import (
    "fmt"
	"bytes"
	"encoding/gob"
)

type TestData struct {
    Content string
}

func Serialize(data Blob, buffer *bytes.Buffer) error {
    fmt.Printf("Buffer addr in Serialize: %p\n", buffer)

	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)


	return err
}

func Deserialize(buffer *bytes.Buffer) ([]byte, error) {

    fmt.Printf("Buffer addr in Deserialize: %p\n", buffer)

	// var b bytes.Buffer
    var b Blob
 
	decoder := gob.NewDecoder(buffer)

    fmt.Println("Bytes to be serialized: ")
    fmt.Println(*buffer)

	for {
        err := decoder.Decode(&b)
		if err != nil {
			fmt.Println("Error in decoder", err)
			break
		}
	}

    fmt.Println("Decoded bytes", string(b.Bytes))

	// decodedData := b.Bytes()
	return b.Bytes, nil
}

package utils

import (
    "bytes"
    "encoding/gob"
)

type TestData struct {
    Num int
    Str string
}

func Serialize(data interface{}) ([]byte, error) {

    var b bytes.Buffer

    encoder := gob.NewEncoder(&b)
    err := encoder.Encode(data)

    serializedData := b.Bytes()
    return serializedData, err
}

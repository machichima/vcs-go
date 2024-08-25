package utils_test

import (
	"log"
	"os"
)

const TestStr = `This is a test
second line`

func CreateTempFile(dir string) (string, error) {
    path := dir + "/test.txt"

    err := os.WriteFile(path, []byte(TestStr), os.ModePerm)
	if err != nil {
		log.Fatal("Error occur while writing temp content to file")
	}

    return path, nil
}


func CreateTempRepo() {
    // TODO: Create a temp repo
}

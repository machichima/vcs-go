package utils

import (
	"fmt"
	"os"
)

type Blob struct {
	Bytes []byte
}

func CheckPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDirs(paths ...string) error {
	for _, path := range paths {
		if err := CreateOneDir(path); err != nil {
			return err
		}
	}
	return nil
}

func GetFiles(path string) ([]string, error) {

	var fileNames []string

	// queue
	queue := make([]string, 0)
	queue = append(queue, path)

	for len(queue) > 0 {

		// pop the first element
		path := queue[0]
		queue = queue[1:]

		dirs, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}
		for _, dir := range dirs {
			if dir.IsDir() {
				fmt.Println(dir.Name())
				queue = append(queue, path+"/"+dir.Name())
			} else {
				fileNames = append(fileNames, path+"/"+dir.Name())
			}
		}
	}

	fmt.Println(fileNames)

	return fileNames, nil
}

func CreateOneDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func FileToStruct(path string) (Blob, error) {
	// Read the file to be serialized
	readBytes, err := os.ReadFile("./test.txt")
	if err != nil {
		return Blob{}, err
	}

	blob := Blob{Bytes: readBytes}

	return blob, nil

}

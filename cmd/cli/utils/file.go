package utils

import (
	"os"
)

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


func CreateOneDir(path string) error {
    return os.MkdirAll(path, os.ModePerm)
}

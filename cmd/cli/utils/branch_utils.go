package utils

import (
	"os"
	"path/filepath"
)

func PointHEADToCommit(commitHash string) error {

    // check the current HEAD branch
	headBytes, err := os.ReadFile(HEADFileName)
	if err != nil {
		return err
	}
	head := string(headBytes)

    // write commit hash to head
    if err := os.WriteFile(filepath.Join(RefsDirName, head), []byte(commitHash), os.ModePerm); err != nil {
        return err
    }

    return nil
}

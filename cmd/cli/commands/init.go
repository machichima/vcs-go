package commands

import (
	"errors"
	"fmt"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

const (
	rootDirName    = ".vgo"
	objectsDirName = ".vgo/objects"
)

func executeInit() error {

	if exists, err := utils.CheckPathExists(rootDirName); err != nil {
		return err
	} else if exists {
		return errors.New("repository already exists")
	}

    if err := utils.CreateDirs(rootDirName, objectsDirName); err != nil {
        return err
    }

    fmt.Println("Initialized empty vcs-go repository")
        
	return nil
}

package commands

import (
	"errors"
	"github.com/machichima/vcs-go/cmd/cli/utils"
)

func executeInit() error {

	if exists, err := utils.CheckPathExists(utils.RootDirName); err != nil {
		return err
	} else if exists {
		return errors.New("repository already exists")
	}

	if err := utils.CreateDirs(utils.RootDirName, utils.ObjectsDirName); err != nil {
		return err
	}

	// fmt.Println("Initialized empty vcs-go repository")

	return nil
}

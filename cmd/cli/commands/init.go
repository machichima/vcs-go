package commands

import (
	"fmt"
	"os"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

func executeInit() error {

	if exists, err := utils.CheckPathExists(utils.RootDirName); err != nil {
		return err
	} else if exists {
        fmt.Println("repository already initialized")
	}

	if err := utils.CreateDirs(utils.RootDirName, utils.ObjectsDirName, utils.RefsDirName); err != nil {
		return err
	}

    // create HEAD dir ref to main branch
    HEADContent := "main"
    if err := os.WriteFile(utils.HEADFileName, []byte(HEADContent), os.ModePerm); err != nil {
        return err
    }

	// fmt.Println("Initialized empty vcs-go repository")

	return nil
}

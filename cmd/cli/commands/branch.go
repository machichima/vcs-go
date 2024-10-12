package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

// create, list, and delete branches. Using branch without any
// arguments and flag will list all branches
func executeBranch(name string, isDelete bool) error {

	// illegal command
	if name == "" && isDelete {
		fmt.Println("Please provide the branch name to delete")
		return nil
	}

	if name != "" {
		// create new branch
		if err := os.WriteFile(
			filepath.Join(utils.RefsDirName, name),
			[]byte{},
			os.ModePerm,
		); err != nil {
			return err
		}

	} else {

		// print all branches and mark the current one
		branchFs, err := os.ReadDir(utils.RefsDirName)
		if err != nil {
			return err
		}

		// read HEAD
		headByte, err := os.ReadFile(utils.HEADFileName)
		if err != nil {
			return err
		}

		for _, fs := range branchFs {
			if fs.Name() == string(headByte) {
				fmt.Printf("*%s\n", fs.Name())
				continue
			}
			fmt.Println(fs.Name())
		}
	}

	if isDelete {
        if _, err := os.ReadFile(filepath.Join(utils.RefsDirName, name)); os.IsNotExist(err) {
            fmt.Printf("the branch %s is not exist\n", name)
            return nil
        } else if err != nil {
            return err
        }

        if err := os.Remove(filepath.Join(utils.RefsDirName, name)); err != nil {
            return err
        }

        fmt.Printf("%s branch deleted", name)
	}

	return nil
}

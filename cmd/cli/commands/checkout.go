package commands

import (
	"fmt"
	"os"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

// Three usecases:
//
// - checkout filename
//   - take the version of the file in the head commit and replace the current file
//
// - checkout COMMIT_ID filename
//   - same as above, but take the version of the specific COMMIT
//
// - checkout COMMIT_ID
//   - same as above, but replace all files in the commit to the current files
//
// - checkout branch_name
//   - change branch (later)
//
// Note: one file only for usage with filename
func executeCheckout(commitHash string, fileName string) error {
	fmt.Println("commit hash: ", commitHash)
	fmt.Println("fileName: ", fileName)

	// Checkout fileName with the HEAD commit version

	// no commit hash provided, use the HEAD
	// get head commit
	if commitHash == "" {
		headbyte, err := os.ReadFile(utils.HEADFileName)
		if err != nil {
			return err
		}
		commitHash = string(headbyte)
	}

    commit, err := utils.ReadCommit(commitHash)
	if err != nil {
		return err
	}

	// read file tree
	fileTree, err := utils.ReadFileTree(commit.FileTree)
	if err != nil {
		return err
	}

	// no fileName provided, checkout all files in the filetree
	// of the commit
	if fileName == "" {

		// loop through all files
		for file, hash := range fileTree.FileToHash {
			// read committed content for the file
			commitByte, err := utils.ReadFileBlobWithSerialize(hash)
			if err != nil {
				return err
			}

			// write committed content to the workspace
			if err := os.WriteFile(file, []byte(commitByte), os.ModePerm); err != nil {
				return err
			}
		}

	} else {
		// filename provided
		hash := fileTree.FileToHash[fileName]
		if hash == "" {
			// fileName does not exist
			return fmt.Errorf("Error finding file %s in the previous commit, %s", fileName, os.ErrNotExist)
		}

		// read committed content for the file
		commitByte, err := utils.ReadFileBlobWithSerialize(hash)
		if err != nil {
			return err
		}

		// write committed content to the workspace
		if err := os.WriteFile(fileName, []byte(commitByte), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

// Three usecases:
//
// - checkout -f filename
//   - take the version of the file in the head commit and replace the current file
//
// - checkout -c COMMIT_ID -f filename1 -f filename2
//   - same as above, but take the version of the specific COMMIT
//
// - checkout -c COMMIT_ID
//   - same as above, but replace all files in the commit to the current files
//
// - checkout -b branch_name
//   - change branch (later)
//
// Note: one file only for usage with filename
func executeCheckout(commitHash string, fileNames []string, branchName string) error {

	// checkout branch
	if branchName != "" {

        if utils.CheckCurrBranch(branchName) {
            fmt.Println("Already in branch ", branchName)
            return nil
        }

        if utils.CheckBranchExist(branchName) {
            // point HEAD to branch
            if err := os.WriteFile(utils.HEADFileName, []byte(branchName), os.ModePerm); err != nil {
                return err
            }
        }

        // clear INDEX if switch
        if err := os.WriteFile(utils.IndexDirName, []byte{}, os.ModePerm); err != nil {
            return err
        }

		return nil
	}

	// no commit hash provided, use the HEAD
	// get head commit
	if commitHash == "" {
		headPathByte, err := os.ReadFile(utils.HEADFileName)
		headByte, err := os.ReadFile(filepath.Join(utils.RefsDirName, string(headPathByte)))

		if err != nil {
			return err
		}
		commitHash = string(headByte)
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
	if len(fileNames) == 0 {

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
		// filenames provided
		for _, f := range fileNames {

			hash := fileTree.FileToHash[f]
			if hash == "" {
				// fileName does not exist
				return fmt.Errorf("error finding file %s in the previous commit", f)
			}

			// read committed content for the file
			commitByte, err := utils.ReadFileBlobWithSerialize(hash)
			if err != nil {
				return err
			}

			// write committed content to the workspace
			if err := os.WriteFile(f, []byte(commitByte), os.ModePerm); err != nil {
				return err
			}
		}
	}

	return nil
}

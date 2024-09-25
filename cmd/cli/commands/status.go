package commands

import (
	"errors"
	"fmt"
	"os"
	"slices"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

// steps:
// 1. check if index file exists (if not, print no staged and return)
// 2. get the bytes content in the index file
// 3. deserialize the bytes content into Index struct
// stages:
// - staged files
//   - modified (in index)
//   - Removed files (not in file tree)
//
// - Modifications Not Staged For Commit (in commit filetree)
//   - modified
//   - deleted
//
// - Untracked Files (not in file tree)
func executeStatus() error {

	index, err := utils.ReadIndexFile()
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	isFirstCommit := false

	commitHashByte, err := os.ReadFile(utils.HEADFileName)
	if err != nil {
		if os.IsNotExist(err) {
			isFirstCommit = true
		} else {
			return err
		}
	}

	var modifiedFiles []string
	var newFiles []string
	var deletedFiles []string
	var stagedModifiedFiles []string
	var stagedNewFiles []string
	var stagedDeletedFiles []string

	// get all files in the workspace
	files, err := utils.GetFiles("./")
	if err != nil {
		return err
	}

	// Go through the filetree of prev commit
    var fileTree utils.Index
	if !isFirstCommit {
		commitHash := string(commitHashByte)

		// get the commit struct by commitHash
		commit, err := utils.ReadCommit(commitHash)
		if err != nil {
			return err
		}

		// Receive the hash of the file with path same as filePath
		fileTreeHash := commit.FileTree
		fileTree, err = utils.ReadFileTree(fileTreeHash)
		if err != nil {
			return err
		}

		for file, _ := range fileTree.FileToHash {
			if !slices.Contains(files, file) {
				// file is deleted since last commit

				if _, ok := index.FileToHash[file]; ok {
					stagedDeletedFiles = append(stagedDeletedFiles, file)
				} else {
					deletedFiles = append(deletedFiles, file)
				}
			}
		}

	}

	// new or modified
	for _, file := range files {

        var fileStatus int = utils.NewFile
        if !isFirstCommit {
            fileStatus, err = utils.CompareFileToFileTree(file, fileTree)
            if err != nil {
                return err
            }
        }

		// file is staged
		if _, ok := index.FileToHash[file]; ok {
			if fileStatus == utils.NewFile {
				stagedNewFiles = append(stagedNewFiles, file)
			} else if fileStatus == utils.ModifiedFile {
				stagedModifiedFiles = append(stagedModifiedFiles, file)
			}
		} else {

			if fileStatus == utils.NewFile {
				newFiles = append(newFiles, file)
			} else if fileStatus == utils.ModifiedFile {
				modifiedFiles = append(modifiedFiles, file)
			}

		}

	}

	fmt.Println("Staged: ")
	for _, file := range stagedNewFiles {
		fmt.Println(file, " (new)")
	}
	for _, file := range stagedModifiedFiles {
		fmt.Println(file, " (modified)")
	}
	for _, file := range stagedDeletedFiles {
		fmt.Println(file, " (deleted)")
	}

	fmt.Println("\nModifications Not Staged For Commit: ")
	for _, file := range newFiles {
		fmt.Println(file, " (new)")
	}
	for _, file := range modifiedFiles {
		fmt.Println(file, " (modified)")
	}
	for _, file := range deletedFiles {
		fmt.Println(file, " (deleted)")
	}

	return nil
}

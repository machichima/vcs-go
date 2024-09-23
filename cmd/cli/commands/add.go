package commands

import (
	// "errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

// steps:
//  1. go through all files, get list of files
//  2. serialize the files
//  3. hash the files to get sha-1 hash
//  4. put files to objects dir (based on hash)
func executeAdd(filePath string) error {

	// check if the file / folder is deleted (following cases)
	// 1) if stage one file and the file does not exist (deleted)
	// 2) if stage one file and the file does not exist (not exist)
	// 3) stage one folder and the folder does not exist (deleted)
	// 4) stage one folder and the folder does not exist (not exist)
	// 5) stage one folder and some files / folders inside it does not exist (deleted)

	// possible methods:
	// array containing all files in the file tree (empty if is first commit)
	// array containing all files in the staged path
	//   (is not dir, one file only) (not contain not exist file)
	// array containing files that does not exist in the currect dir
	// if file does not exist exists in committed files, go through delete file
	//   process
	// Go through the staging process for staged files

    // standarize the format of the filePath
    filePath = filepath.Join(filePath)

	var deletedFiles []string
	stat, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// file does not exist in currect directory
			deletedFiles = append(deletedFiles, filePath)
		} else {
			return err
		}
	}

	// Go through whole files in the dir
	// create files arr (for staged files)
	var files []string
	// if not the path not exist
	if len(deletedFiles) == 0 {
		if stat.IsDir() {
			files, err = utils.GetFiles(filePath)
			if err != nil {
				return err
			}
		} else {
			files = append(files, filePath)
		}
	}

	// Go through whole files in the filetree
	// (if not first commit)
	isFirstCommit := false

	commitHashByte, err := os.ReadFile(utils.HEADFileName)
	if err != nil {
		if os.IsNotExist(err) {
			isFirstCommit = true
		} else {
			return err
		}
	}

	// Get the committed filetree
	var fileTree utils.Index

	if !isFirstCommit {
		commitHash := string(commitHashByte)
		// get the commit struct by commitHash
		commit, err := utils.ReadCommit(commitHash)
		if err != nil {
			return err
		}
		// Get file tree of previous commit
		fileTreeHash := commit.FileTree
		fileTree, err = utils.ReadFileTree(fileTreeHash)
		if err != nil {
			return err
		}

		// detect deleted files in the dir (if the path is dir)
		if len(deletedFiles) == 0 && stat.IsDir() {
			for file, _ := range fileTree.FileToHash {
				if !slices.Contains(files, file) {
					// file deleted
					deletedFiles = append(deletedFiles, file)
				}
			}
		}

	}

	fmt.Println("current dir:")
	fmt.Println(files)

	fmt.Println("deleted files:")
	fmt.Println(deletedFiles)

	fmt.Println("committed filetree")
	fmt.Println(fileTree)

	// adding new files to objects
	var isAddNewFile bool = false

	// handle staging files in the current dir (not deleted files)
	for _, file := range files {

		if !isFirstCommit {
			fileStatus, err := utils.CompareFileToFileTree(file, fileTree)
			if err != nil {
				return err
			}
			// only the new or modified files will pass through
			// following code for saving files
			if !(fileStatus == utils.NewFile || fileStatus == utils.ModifiedFile) {
				continue
			}
		}

		isNewFile, err := utils.WriteFileBlobWithSerialize(file)
		if err != nil {
			return err
		}

		isAddNewFile = isAddNewFile || isNewFile
	}

	for _, file := range deletedFiles {
		if _, err := utils.SaveFileByHash(file, "", []byte("")); err != nil {
			return err
		}
	}

	if isAddNewFile {
		fmt.Println("Files added successfully")
	} else {
		fmt.Println("No new files added")
	}

	return nil
}

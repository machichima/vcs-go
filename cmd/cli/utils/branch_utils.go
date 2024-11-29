package utils

import (
	"errors"
	"fmt"
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

// Check if branch is same as current branch.
// If yes, return true
func CheckCurrBranch(branchName string) bool {

	currBranchByte, _ := os.ReadFile(HEADFileName)
	if branchName == string(currBranchByte) {
		return true
	}

	return false

}

// Check whether branch exist
func CheckBranchExist(branchName string) bool {

	branchFs, err := os.ReadDir(RefsDirName)
	if err != nil {
		return false
	}

	var isBranchExist bool = false
	for _, fs := range branchFs {
		if fs.Name() == branchName {
			isBranchExist = true
			break
		}
	}

	return isBranchExist

}

func GetPastCommits(branchName string) ([]string, error) {
	branchFs, err := os.ReadDir(RefsDirName)
	if err != nil {
		return []string{}, err
	}

	var parent string

	// get the latest commit of the branch
	for _, fs := range branchFs {
		if fs.Name() == branchName {
			commitStrBytes, err := os.ReadFile(filepath.Join(RefsDirName, fs.Name()))
			if err != nil {
				return []string{}, err
			}
			// commitStrList = append(commitStrList, string(commitStrBytes))
			parent = string(commitStrBytes)
			break
		}
	}

	// if branch not exist
	if parent == "" {
		return []string{}, errors.New(
			fmt.Sprintf("Branch %s not exist", branchName),
		)
	}

	commitStrList := make([]string, 0)
	for parent != "" {
		// add parent commit to the list
		commitStrList = append(commitStrList, parent)
		commit, err := ReadCommit(parent)
		if err != nil {
			return []string{}, err
		}
		parent = commit.ParentCommit
	}

	return commitStrList, nil
}

func MatchCommitInList(targetCommitHash string, commitList []string)

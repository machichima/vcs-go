package commands_test

import (
	"testing"
)

const CommandPrefix string = "go run ../main.go"
const dir string = "./testfolder/"

func TestAddOneFile(t *testing.T) {
	// init test folder (init repo)
	// create temp dir
	testCaseFile := "./testCases/addOneFile.txt"

	RunTestCases(dir, testCaseFile, t)
}

func TestAddOneDir(t *testing.T) {
	// init test folder (init repo)
	// create temp dir
	testCaseFile := "./testCases/addOneDir.txt"

	RunTestCases(dir, testCaseFile, t)
}

func TestCommit(t *testing.T) {
	// init test folder (init repo)
	// create temp dir
	testCaseFile := "./testCases/commit.txt"

	RunTestCases(dir, testCaseFile, t)
}

func TestRm(t *testing.T) {
	// init test folder (init repo)
	// create temp dir
	testCaseFile := "./testCases/rm.txt"

	RunTestCases(dir, testCaseFile, t)
}

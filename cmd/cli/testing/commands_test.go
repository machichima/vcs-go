package commands_test

import (
	"testing"
)

const CommandPrefix string = "go run ../main.go"
const dir string = "./testfolder/"


func TestInit(t *testing.T) {
	testCaseFile := "./testCases/init.txt"

	RunTestCases(dir, testCaseFile, t)
}

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

func TestCheckoutNoCommitFileExist(t *testing.T) {
	// init test folder (init repo)
	// create temp dir
	testCaseFile := "./testCases/checkoutNoCommitFileExist.txt"

	RunTestCases(dir, testCaseFile, t)

}

func TestCheckoutWithCommitFileExist(t *testing.T) {
	// init test folder (init repo)
	// create temp dir
	testCaseFile := "./testCases/checkoutWithCommitFileExist.txt"

	RunTestCases(dir, testCaseFile, t)
}

func TestCheckoutWithCommitNoFile(t *testing.T) {
	// init test folder (init repo)
	// create temp dir
	testCaseFile := "./testCases/checkoutWithCommitNoFile.txt"

	RunTestCases(dir, testCaseFile, t)
}

func TestCheckoutNoCommitMultiFiles(t *testing.T) {
	testCaseFile := "./testCases/checkoutNoCommitMultiFiles.txt"

	RunTestCases(dir, testCaseFile, t)
}

func TestBranchOpera(t *testing.T) {
	testCaseFile := "./testCases/createAndDeleteBranch.txt"

	RunTestCases(dir, testCaseFile, t)
}

func TestSwitchBranch(t *testing.T) {
	testCaseFile := "./testCases/switchBranch.txt"

	RunTestCases(dir, testCaseFile, t)
}

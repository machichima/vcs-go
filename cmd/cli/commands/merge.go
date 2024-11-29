package commands

import (
	"fmt"
	"os"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

func executeMerge(targetBranch string) error {

    if utils.CheckCurrBranch(targetBranch) {
        fmt.Println("Cannot merge with current branch")
    }
    if !utils.CheckBranchExist(targetBranch) {
        fmt.Println("target branch not exist")
    }

    currBranchByte, _ := os.ReadFile(utils.HEADFileName)
    currBranch := string(currBranchByte)

    currBranchHis := []string{}
    targetBranchHis := []string{}

    currBranchHisQue := []string{currBranch}
    targetBranchHisQue := []string{targetBranch}

    utils.GetPastCommits
    // get all past commits of source and target branch
    // consider second parent for merged commit

    // loop curr branch to the end, memorize all commit hashes
    // loop target branch, see whether the commit is in the curr branch hashes
    

    // detect different cases

    return nil
}


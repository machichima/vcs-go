package commands

import (
	"fmt"
	"os"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

func executeLog() error {

	// read the commit with hash in HEAD
	headBytes, err := os.ReadFile(utils.HEADFileName)
	if err != nil {
		return err
	}
	headCommitHash := string(headBytes)

	// queue for going through whole commit history
	queue := make([]string, 0)
	queue = append(queue, headCommitHash)

	for len(queue) > 0 {
		hash := queue[0]
		queue = queue[1:]

		// get the commit file by the hash
        commit, err := utils.ReadCommit(hash)
        if err != nil {
            return err
        }

        fmt.Println(commit)

	}

    // TODO: format the output of the log

	return nil
}

package commands

import (
	"fmt"
    "errors"
    "os"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

// format: vgo commit -m "message"
// 1. check if message if provided
// 2. check if index file exists and if it is empty
// 3. get the content of the index (file and hash)
// 4. create file tree object (index object)
// 5. create commit object
// - contain commit message, file tree hash, parent commit hash
// 6. save commit and file tree objects to objects dir
// 7. HEAD file points to the commit object
func executeCommit(msg string) error {

    if msg == "" {
        return fmt.Errorf("Please provide the commit messages")
    }

    index, err := utils.ReadIndexFile()
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println("No staged changes")
            return nil
        } else {
            return err
        }
    }

    if len(index.FileToHash) < 1 {
        fmt.Println("No staged changes")
        return nil
    }

    hash, err := utils.WriteFileTree(index)
    if err != nil {
        return err
    }

    // check if HEAD files exists and retrieve the previous commit hash
    headCommitHash, err := os.ReadFile(utils.HEADFileName)
    if err != nil{
        if !os.IsNotExist(err) {
            return err
        }
    }

    commit := utils.Commit{
        Message: msg,
        FileTree: hash,
        ParentCommit: string(headCommitHash),
    }

    commitHash, err := utils.WriteCommit(commit)
    if err != nil {
        return err
    }

    // update HEAD file
    if err := os.WriteFile(utils.HEADFileName, []byte(commitHash), os.ModePerm); err != nil {
        return err
    }

    return nil
}

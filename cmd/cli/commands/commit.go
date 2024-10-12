package commands

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

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
        fmt.Println("Please provide the commit messages")
        return nil
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

	headByte, err := os.ReadFile(utils.HEADFileName)
	headCommitHash, err := os.ReadFile(filepath.Join(utils.RefsDirName, string(headByte)))
    // check if HEAD files exists and retrieve the previous commit hash
    if err != nil{
        if !os.IsNotExist(err) {
            return err
        }
    }

    // if headCommitHash != nil, read the previous filetree
    // and merge with the current one
    newFileTree := index
    if string(headCommitHash) != "" {
        prevCommit, err := utils.ReadCommit(string(headCommitHash))
        if err != nil {
            return err
        }
        prevFileTree, err := utils.ReadFileTree(prevCommit.FileTree)
        if err != nil {
            return err
        }

        newFileTree = utils.MergeIndexAndFileTree(index, prevFileTree)
    }

    // merge index filetree with the previous committed files
    hash, err := utils.WriteFileTree(newFileTree)
    if err != nil {
        return err
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
    if err := utils.PointHEADToCommit(commitHash); err != nil {
        return err
    }

    // clear the index file that is committed
    os.WriteFile(utils.IndexDirName, []byte(""), os.ModePerm)

    return nil
}

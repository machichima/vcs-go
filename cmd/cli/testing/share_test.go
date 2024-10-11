package commands_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/machichima/vcs-go/cmd/cli/utils"
)

var FileStruct []string = []string{"test_1.txt", "test_2.txt", "test/test_3.txt", "test/test_4.txt"}

var MainPath string = "../../main.go"

func RunTestCases(dir string, testCaseFile string, t *testing.T) {

	if err := CreateTempRepo(dir); err != nil {
		t.Error(err)
	}

	// read test cases
	commands, expectedOutputs, err := ReadTestCases(testCaseFile)
	if err != nil {
		t.Errorf("Err when reading test cases file: %s", err)
	}

    var commitHashs []string
	for i, c := range commands {

        fmt.Printf("test command: %s\n", c)

        // detect {HASH} tag
        re := regexp.MustCompile(`{(.*?)}`)
        loc := re.FindStringIndex(c)
        if loc != nil && len(commitHashs) > 0 {  // {HASH} detected
            hashTag := strings.Trim(string(c[loc[0]:loc[1]]), "{}")
            hashIndex, err := strconv.Atoi(strings.Split(hashTag, "_")[1])
            if err != nil {
                t.Error(err)
            }
            c = strings.Join([]string{c[:loc[0]], commitHashs[hashIndex], c[loc[1]:]}, "")
        }

        output, currCommitHashs, err := ExecCommand(c, dir, strings.ContainsAny(c, "!"))
		output = strings.Trim(output, "\n")
		if err != nil {
			t.Errorf("Err exec command: %s", err)
		}

		if output != expectedOutputs[i] {
			t.Errorf("Output mismatch for command %s\n", c)
			t.Errorf("Expected output: \n%s\n", expectedOutputs[i])
			t.Errorf("Output get: \n%s\n", output)
		}

        // update commitHashs if currCommitHashs is not empty
        if len(currCommitHashs) > 0 {
            commitHashs = currCommitHashs
        }

	}

}

// create temp repo with temp files, the content\
// of the temp file is the filename
func CreateTempRepo(dir string) error {

	for _, file := range FileStruct {
		fullPath := filepath.Join(dir, file)

		if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
			return fmt.Errorf("Err when creating files' dir: %s", err)
		}

		if err := os.WriteFile(fullPath, []byte(file), os.ModePerm); err != nil {
			return fmt.Errorf("Err when creating files: %s", err)
		}
	}

	// remove .vgo file if exist
	vgoPath := filepath.Join(dir, utils.RootDirName)
	if _, err := os.Stat(vgoPath); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		os.RemoveAll(vgoPath)
		fmt.Println("Remove .vgo")
		return nil
	}

	return nil
}

// return commands and expectedOutputs from the testcase file
//
// Note that the expectedOutputs are squeeze into one line and
// emit any linebreak and spaces in between
func ReadTestCases(testCasesFile string) ([]string, []string, error) {

	testCasesByte, err := os.ReadFile(testCasesFile)
	if err != nil {
		return []string{}, []string{}, fmt.Errorf("Err reading test cases: %s", err)
	}
	testCases := string(testCasesByte)

	commandBlocks := strings.Split(testCases, "> ")

	var commands []string
	var expectedOutputs []string
	// fmt.Println(len(commandBlocks))
	// fmt.Println(strings.Join(commandBlocks, ", "))
	for _, commandBlock := range commandBlocks {
		if len(commandBlock) < 1 {
			// empty block
			continue
		}
		commandRes := strings.Split(commandBlock, "<<<")

		commands = append(commands, strings.Trim(commandRes[0], "\n"))

		// deal with output format (squeeze the output into onel line)
		// and emit any spaces on the both end of the sentence

		outSentences := strings.Split(commandRes[1], "\n")
		for i, _ := range outSentences {
			outSentences[i] = strings.Trim(outSentences[i], " ")
		}

		output := strings.Join(outSentences, "")

		expectedOutputs = append(expectedOutputs, output)
	}

	return commands, expectedOutputs, nil
}


// input "command" to execute and "dir" for
// the directory to exec the command. If the command will produce
// commit hash, store them in commitHashs and return.
//
// return output, commitHashs, and err. If not commit hash exists, 
// return empty array {} for commitHashs
func ExecCommand(command string, dir string, isTermCmd bool) (output string, commitHashs []string, err error) {
	// cmd := exec.Command("go", "run", "main.go", command)
	// separate commands
	re := regexp.MustCompile(`"(.*?)"|\S+`)
	cmdArr := re.FindAllString(command, -1)

	// Trim double quote "
	for i, _ := range cmdArr {
		cmdArr[i] = strings.Trim(cmdArr[i], `"`)
	}

    cmdArgs := []string{}
    var cmdName string
	if isTermCmd {
        cmdArr[0] = strings.Trim(cmdArr[0], "!")
        cmdName = "/bin/sh"

        cmdStr := strings.Join(cmdArr, " ")

        cmdArgs = []string{"-c"}
        cmdArgs = append(cmdArgs, cmdStr)

	} else {
        cmdName = "go"
        cmdArgs = []string{"run", MainPath}
        cmdArgs = append(cmdArgs, cmdArr...)
    }
    cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Dir = dir

	outputByte, err := cmd.CombinedOutput()
	if err != nil {
		return string(outputByte), []string{}, err
	}

	outSentences := strings.Split(string(outputByte), "\n")
	for i, _ := range outSentences {
		if isCommitCmd := strings.Contains(outSentences[i], "commit"); isCommitCmd {

			if commitMsgArr := strings.Split(outSentences[i], ": "); len(commitMsgArr) > 1 {
				// deal with "commit: hash" for commit cmd

                commitHashs = append(commitHashs, commitMsgArr[1])

				commitMsgArr[1] = "{HASH}"
				outSentences[i] = strings.Join(commitMsgArr, ": ")
			}
		}
		outSentences[i] = strings.Trim(outSentences[i], " ")
	}

	output = strings.Join(outSentences, "")

	return output, commitHashs, nil
}

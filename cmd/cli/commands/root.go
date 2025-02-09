package commands

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v2" // imports as package "cli"
)

// TODO: take out cli tool and just receive args

func ExecuteCommands() {
	app := &cli.App{
		Name:  "vgo",
		Usage: "vcs-go (vgo) is a simple version control system",
		Commands: []*cli.Command{
			InitSubCmd,
			AddSubCmd,
			StatusSubCmd,
			CommitSubCmd,
			LogSubCmd,
			RmSubCmd,
			CheckoutSubCmd,
            BranchSubCmd,
            MergeSubCmd,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var InitSubCmd = &cli.Command{
	Name:  "init",
	Usage: "init initializes a new vcs-go repository",
	Action: func(c *cli.Context) error {
		return executeInit()
	},
}

var AddSubCmd = &cli.Command{
	Name:  "add",
	Usage: "stage changed files",
	Action: func(c *cli.Context) error {
		filePath := c.Args().First()
		return executeAdd(filePath)
	},
}

var StatusSubCmd = &cli.Command{
	Name:  "status",
	Usage: "show the staged files",
	Action: func(c *cli.Context) error {
		return executeStatus()
	},
}

var CommitSubCmd = &cli.Command{
	Name:  "commit",
	Usage: "commit the staged files",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "message",
			Aliases: []string{"m"},
			Usage:   "commit message",
		},
	},
	Action: func(c *cli.Context) error {
		msg := c.String("message")
		return executeCommit(msg)
	},
}

var LogSubCmd = &cli.Command{
	Name:  "log",
	Usage: "show the commit history",
	Action: func(c *cli.Context) error {
		return executeLog()
	},
}

var RmSubCmd = &cli.Command{
	Name:  "rm",
	Usage: "unstaged the file or dir",
	Action: func(c *cli.Context) error {
		filePath := c.Args().First()
		return executeRm(filePath)
	},
}


var commitHash string
var files []string

var CheckoutSubCmd = &cli.Command{
	Name:  "checkout",
	Usage: "checkout file",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "commit",
            Aliases: []string{"c"},
			Usage:       "commit hash to checkout",
			Destination: &commitHash,
		},
        &cli.StringSliceFlag{
			Name:        "file",
            Aliases: []string{"f"},
			Usage:       "files to checkout",
        },
        &cli.StringFlag{
			Name:        "branch",
            Aliases: []string{"b"},
			Usage:       "checkout branch",
        },
	},
	Action: func(c *cli.Context) error {
        return executeCheckout(commitHash, c.StringSlice("file"), c.String("branch"))
	},
}

var BranchSubCmd = &cli.Command{
	Name:  "branch",
	Usage: "create new branches",
	Flags: []cli.Flag{
        &cli.BoolFlag{
			Name:        "delete",
            Aliases: []string{"d"},
			Usage:       "delete branch",
        },
	},
	Action: func(c *cli.Context) error {
        return executeBranch(c.Args().First(), c.Bool("delete"))
	},
}

var MergeSubCmd = &cli.Command{
	Name:  "merge",
	Usage: "merge current branch to target branch",
	Action: func(c *cli.Context) error {
        if c.Args().Len() < 1 {
            return errors.New("Please provide target branch to merge")
        }
        if c.Args().Len() > 1 {
            return errors.New("Too many arguments. Only one target branch needed")
        }
		targetBranch := c.Args().First()
		return executeMerge(targetBranch)
	},
}

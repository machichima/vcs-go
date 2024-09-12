package commands

import (
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
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

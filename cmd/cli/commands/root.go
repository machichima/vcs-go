package commands

import (
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
)

func ExecuteCommands() {
	app := &cli.App{
		Name:  "vgo",
		Usage: "vcs-go (vgo) is a simple version control system",
		Commands: []*cli.Command{
			InitSubCmd,
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

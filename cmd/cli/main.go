package main

import (
	"log"

	"github.com/machichima/vcs-go/cmd/cli/commands"
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
	commands.ExecuteCommands()
}

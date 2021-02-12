package main

import (
	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("subcmd", "config")
	x.Summary = `sample command with config subcommand`
	x.Description = `
		The subcmd command simple demonstrates what is possible by
		including the config.go file within your own commands.`

}

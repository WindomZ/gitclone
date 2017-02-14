package gitclone

import (
	"github.com/WindomZ/cli-mate"
	"github.com/urfave/cli"
)

var CloneCommand = cli_mate.Command{
	Command: cli.Command{
		Name:    "clone",
		Aliases: []string{"c"},
		Usage:   "same as `git clone`",
	},
	Action: func(c *cli_mate.Context) cli_mate.ExitCoder {
		c.Println("In development...")
		return nil
	},
}

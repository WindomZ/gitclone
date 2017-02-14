package gitclone

import (
	"github.com/WindomZ/cli-mate"
	"github.com/urfave/cli"
)

var ListCommand = cli_mate.Command{
	Command: cli.Command{
		Name:    "list",
		Aliases: []string{"ls", "l"},
		Usage:   "prints a list of files in the current directory",
	},
	Action: func(c *cli_mate.Context) cli_mate.ExitCoder {
		c.Println("In development...")
		return nil
	},
}

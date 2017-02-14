package gitclone

import (
	"github.com/WindomZ/cli-mate"
	"github.com/urfave/cli"
	"os"
)

var MkdirCommand = cli_mate.Command{
	Command: cli.Command{
		Name:    "mkdir",
		Aliases: []string{"dir", "d"},
		Usage:   "create a default directory",
	},
	Action: func(c *cli_mate.Context) cli_mate.ExitCoder {
		if !ExistFile(DEFAULT_DIR_NAME) {
			if err := os.Mkdir(DEFAULT_DIR_NAME, os.ModePerm); err != nil {
				return cli_mate.NewExitError(err)
			}
			c.Println("Create a default directory")
			return nil
		}
		c.Println("Directory already exists")
		return nil
	},
}

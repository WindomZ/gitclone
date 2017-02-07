package cmd

import (
	"github.com/urfave/cli"
	"os"
)

const DEFAULT_DIR_NAME string = "gitclone"

var InitFlagAction = &FlagAction{
	Flag: cli.BoolFlag{
		Name:  "init, i",
		Usage: "Initialize the default directory",
	},
	FlagName: "init",
	Action: func(c *cli.Context, f *FlagAction) (bool, string, error) {
		if !c.Bool(f.FlagName) {
			return false, "", nil
		}
		if !ExistFile(DEFAULT_DIR_NAME) {
			if err := os.Mkdir(DEFAULT_DIR_NAME, os.ModePerm); err != nil {
				return true, "", err
			}
			return true, "Create a default directory", nil
		}
		return true, "Directory already exists", nil
	},
}

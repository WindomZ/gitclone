package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var RootFlagAction = &FlagAction{
	Flag:     nil,
	FlagName: "",
	Action: func(c *cli.Context, f *FlagAction) (bool, *cli.ExitError) {
		if len(c.Args()) == 0 {
			return true, cli.NewExitError("No args!", 86) // error
		}
		println(fmt.Sprintf("root-value: %#v", c.Args())) // log
		return true, nil
	},
}

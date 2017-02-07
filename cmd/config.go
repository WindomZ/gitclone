package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var ConfigFlagAction = &FlagAction{
	Flag: cli.StringFlag{
		Name:  "config, c",
		Usage: "Load configuration from `FILE`",
	},
	FlagName: "config",
	Action: func(c *cli.Context, f *FlagAction) (bool, error) {
		value := c.String(f.FlagName)
		if value == "" {
			return false, nil
		}
		println(fmt.Sprintf("config-value: %#v", value)) // log
		return false, nil
	},
}

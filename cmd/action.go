package cmd

import "github.com/urfave/cli"

type FlagAction struct {
	Flag     cli.Flag
	FlagName string
	Action   func(c *cli.Context, f *FlagAction) (interrupt bool, err error)
}

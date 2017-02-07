package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
)

type Cmd struct {
	app         *cli.App
	flagActions []*FlagAction
}

func newCmd(app *cli.App) *Cmd {
	return &Cmd{
		app: app,
	}
}

func NewCmd() *Cmd {
	return newCmd(cli.NewApp())
}

func NewSimpleCmd(name, usage, version string) *Cmd {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.Version = version
	return newCmd(app)
}

func (cmd *Cmd) GetApp() *cli.App {
	return cmd.app
}

func (cmd *Cmd) Run(arguments []string) error {
	if arguments != nil && len(arguments) > 1 {
		//println(fmt.Sprintf("%#v", arguments)) // log
		cmd.app.Action = func(c *cli.Context) error {
			for _, f := range cmd.flagActions {
				if interrupt, content, err := f.Action(c, f); interrupt {
					if err != nil {
						return errors.New(err.Error())
					}
					fmt.Println(content)
					fmt.Println("Finish!")
					return nil
				}
			}
			return nil
		}
	}
	return cmd.app.Run(arguments)
}

func (cmd *Cmd) AddFlagAction(flagAction *FlagAction) {
	cmd.flagActions = append(cmd.flagActions, flagAction)
	if flagAction.Flag != nil {
		cmd.app.Flags = append(cmd.app.Flags, flagAction.Flag)
	}
}

func (cmd *Cmd) AddFlagActions(flagActions []*FlagAction) {
	for _, flagAction := range flagActions {
		cmd.AddFlagAction(flagAction)
	}
}

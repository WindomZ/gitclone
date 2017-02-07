package main

import (
	"os"

	"github.com/WindomZ/gitclone/cmd"
)

func main() {
	_cmd := cmd.NewSimpleCmd(
		"gitclone",
		"A cli tool, git clone repository in the `go get` style.",
		"0.0.2",
	)
	_cmd.AddFlagActions([]*cmd.FlagAction{
		cmd.ConfigFlagAction, // config
		cmd.InitFlagAction,   // init
		cmd.RootFlagAction,   // root
	})
	_cmd.Run(os.Args)
}
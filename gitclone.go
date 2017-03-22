package main

import (
	"github.com/WindomZ/gitclone/gitclone"
	"github.com/WindomZ/go-commander"
)

func main() {
	// init
	commander.Program.
		Command("gitclone").
		Description("A cli tool, git clone repository in the `go get` style.").
		Version("0.3.0")

	// gitclone list
	commander.Program.
		Command("list").
		Description("prints a list of repositories witch in the current directory").
		Action(gitclone.ListAction)

	// gitclone
	commander.Program.
		LineArgument("<repo>").
		Description("git clone repository in the 'go get' style").
		Action(gitclone.RootAction)

	commander.Program.Parse()
}

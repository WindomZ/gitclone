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

	// gitclone
	commander.Program.
		LineArgument("<repo>").
		Action(gitclone.RootAction)

	// gitclone clone <repo>
	commander.Program.
		Command("clone <repo>").
		Action(gitclone.CloneAction)

	// gitclone list
	commander.Program.
		Command("list").
		Action(gitclone.ListAction())

	commander.Program.Parse()
}

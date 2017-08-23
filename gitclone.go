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
		Version("0.7.0")

	// gitclone list
	commander.Program.
		Command("list").
		Aliases([]string{"ls"}).
		Description("prints a list of repositories witch in the current directory").
		Action(gitclone.ListAction)

	// gitclone search
	commander.Program.
		Command("search <key>").
		Description("search repositories witch in current directory").
		Action(gitclone.SearchAction)

	// gitclone link
	commander.Program.
		Command("link <filepath>").
		Description("`gitclone` a repository from the directory to current directory").
		Action(gitclone.LinkAction)

	// gitclone
	commander.Program.
		Command("<repo>").
		Description("git clone repository in the 'go get' style").
		Action(gitclone.RootAction).
		Option("--depth=<depth>")

	commander.Program.Parse()
}

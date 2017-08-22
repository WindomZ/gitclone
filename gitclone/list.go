package gitclone

import "fmt"

// ListAction the command action 'list|ls' implementation
var ListAction = func() error {
	dirs := getGitRepoDirList(getCurrentDirectory())

	if len(dirs) == 0 {
		fmt.Println("Not found any git repository!")
		return nil
	}

	fmt.Println("Found the following git repositories:")
	for i, dir := range dirs {
		fmt.Printf("%4d >>> %s\n", i, dir)
	}
	return nil
}

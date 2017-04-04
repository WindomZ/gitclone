package gitclone

import (
	"fmt"
	"github.com/WindomZ/go-commander"
	"strings"
)

var SearchAction = func(c commander.Context) error {
	repo := strings.ToLower(c.MustString("<key>"))

	dirs := getGitRepoDirList(getCurrentDirectory())

	list := make([]string, 0, len(dirs))
	for _, dir := range dirs {
		if strings.Contains(strings.ToLower(dir), repo) {
			list = append(list, dir)
		}
	}

	if len(list) == 0 {
		fmt.Println("Not found any matching git repository!")
		return nil
	}

	fmt.Println("Found the following matching git repositories:")
	for i, dir := range list {
		fmt.Printf("%4d >>> %s\n", i, dir)
	}
	return nil
}

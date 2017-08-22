package gitclone

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/WindomZ/go-commander"
	"github.com/whilp/git-urls"
)

// RootAction the root command action implementation
var RootAction = func(c commander.Context) error {
	content, err := rootAction(c.MustString("<repo>"))
	if err != nil {
		commander.Program.ShowHelpMessage()
		return err
	}
	fmt.Println(content)
	return nil
}

func rootAction(repo string) (string, error) {
	if !validGitAddress(repo) {
		return "", errors.New(fmt.Sprintf("repository '%v' does not exist", repo))
	}
	u, err := giturls.Parse(repo)
	if err != nil {
		return "", err
	}
	f_dir := strings.Replace(u.Host+u.Path, ".git", "", -1)

	var out string
	if existFile(path.Join(f_dir, ".git")) {
		if err = execChdir(f_dir); err == nil {
			out, err = execCommand("git", "pull")
		}
	} else {
		out, err = execCommand("git", "clone", repo, f_dir)
	}
	return out, err
}

package gitclone

import (
	"fmt"
	"path"
	"strings"

	"os"

	"github.com/WindomZ/go-commander"
	"github.com/whilp/git-urls"
)

// RootAction the root command action implementation
var RootAction = func(c commander.Context) error {
	content, err := rootAction(c.MustString("<repo>"), c.MustInt("--depth"))
	if err != nil {
		commander.Program.ShowHelpMessage()
		return err
	}

	fmt.Println(content)

	return nil
}

func rootAction(repo string, depth int) (string, error) {
	if !validGitAddress(repo) {
		return "", fmt.Errorf("repository '%v' does not exist", repo)
	}
	u, err := giturls.Parse(repo)
	if err != nil {
		return "", err
	}
	fDir := strings.Replace(u.Host+u.Path, ".git", "", -1)

	// split file paths with depth
	if depth >= 1 {
		fPaths := strings.Split(fDir, string(os.PathSeparator))
		if len(fPaths) > depth {
			fDir = path.Join(fPaths[len(fPaths)-depth:]...)
		}
	}

	var out string
	if existFile(path.Join(fDir, ".git")) {
		if err = execChdir(fDir); err == nil {
			out, err = execCommand("git", "pull")
		}
	} else {
		out, err = execCommand("git", "clone", repo, fDir)
	}
	return out, err
}

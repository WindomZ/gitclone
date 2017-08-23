package gitclone

import (
	"fmt"
	"path"

	"github.com/WindomZ/go-commander"
)

// LinkAction the command action 'link' implementation
var LinkAction = func(c commander.Context) error {
	filepath := c.MustString("<filepath>")
	if !isGitRepoDir(filepath) {
		fmt.Println("Not found any git repository!")
		return nil
	}
	fmt.Println("link:", filepath)
	content, err := linkAction(filepath)
	if len(content) != 0 {
		fmt.Println(content)
	}
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func linkAction(fDir string) (out string, err error) {
	if fDir == getCurrentDirectory() {
		return execCommand("git", "pull")
	}
	fPath := path.Base(fDir)
	if existFile(path.Base(fPath)) {
	} else if _, err = execCommand("cp", "-R", fDir, "./"); err != nil {
		return
	}
	fDir = path.Join(getCurrentDirectory(), fPath)
	if err = execChdir(fDir); err != nil {
		return
	} else if out, err = execCommand("git", "pull"); err != nil {
		return
	}
	return
}

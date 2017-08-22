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

func linkAction(f_dir string) (out string, err error) {
	if f_dir == getCurrentDirectory() {
		return execCommand("git", "pull")
	}
	f_path := path.Base(f_dir)
	if existFile(path.Base(f_path)) {
	} else if _, err = execCommand("cp", "-R", f_dir, "./"); err != nil {
		return
	}
	f_dir = path.Join(getCurrentDirectory(), f_path)
	if err = execChdir(f_dir); err != nil {
		return
	} else if out, err = execCommand("git", "pull"); err != nil {
		return
	}
	return
}

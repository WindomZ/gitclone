package gitclone

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/WindomZ/go-commander"
	"github.com/whilp/git-urls"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

var RootAction = func(c commander.Context) error {
	content, err := rootAction(c.GetString("<repo>"))
	if err != nil {
		commander.Program.ShowHelpMessage()
		return err
	}
	fmt.Println(content)
	return nil
}

func rootAction(repo string) (string, error) {
	if !ValidGitAddress(repo) {
		return "", errors.New(fmt.Sprintf("repository '%v' does not exist", repo))
	}
	u, err := giturls.Parse(repo)
	if err != nil {
		return "", err
	}
	f_dir := strings.Replace(u.Host+u.Path, ".git", "", -1)

	var out string
	if existFile(path.Join(f_dir, ".git")) {
		out, err = execCommand("git", "pull", f_dir)
	} else {
		out, err = execCommand("git", "clone", repo, f_dir)
	}
	return out, err
}

func ValidGitAddress(repo string) bool {
	matched, err := regexp.MatchString(
		`((git|ssh|http(s)?)|(git@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(\.git)(/)?`,
		repo)
	return matched && err == nil
}

func execCommand(commandName string, params ...string) (string, error) {
	cmd := exec.Command(commandName, params...)

	stdout, stderr, err := pipeCommand(cmd)
	//out := strings.TrimSpace(strings.Join([]string{stdout, stderr}, ""))
	out := strings.TrimSpace(stdout + stderr)
	if err != nil {
		return out, err
	}

	return out, nil
}

func pipeCommand(cmd *exec.Cmd) (string, string, error) {
	var output bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout, cmd.Stderr = &output, &stderr

	if err := cmd.Start(); err != nil {
		return output.String(), stderr.String(), err
	}
	cmd.Wait()

	return output.String(), stderr.String(), nil
}

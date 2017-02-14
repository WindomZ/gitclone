package gitclone

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/WindomZ/cli-mate"
	"github.com/whilp/git-urls"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

const DEFAULT_DIR_NAME string = "GitClones"

var RootAction cli_mate.CommandAction = func(c *cli_mate.Context) cli_mate.ExitCoder {
	if len(c.Args()) == 0 {
		c.ShowAppHelp()
		//return cli_mate.NewExitStringError("Missing command arguments!")
		return nil
	}
	content, err := rootAction(c.Args().Get(0))
	if err != nil {
		return cli_mate.NewExitError(err)
	}
	c.Println(content)
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
	if ExistFile(DEFAULT_DIR_NAME) {
		f_dir = path.Join(DEFAULT_DIR_NAME, f_dir)
	}

	var out string
	if ExistFile(path.Join(f_dir, ".git")) {
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

	//println(strings.Join(cmd.Args, " "))

	stdout, stderr, err := pipeCommand(cmd)
	//out := strings.TrimSpace(strings.Join([]string{stdout, stderr}, ""))
	out := strings.TrimSpace(stdout + stderr)
	if err != nil {
		return out, err
	}

	//println(fmt.Sprintf("out: %v", stdout))
	//println(fmt.Sprintf("err: %v", stderr))
	//println(fmt.Sprintf("sss: %v", out))

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

package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/urfave/cli"
	"github.com/whilp/git-urls"
)

var RootFlagAction = &FlagAction{
	Flag:     nil,
	FlagName: "",
	Action: func(c *cli.Context, f *FlagAction) (bool, string, error) {
		if len(c.Args()) == 0 {
			// print usage
			return true, "", errors.New("No args!") // error
		}
		repo := c.Args().Get(0)
		if !ValidGitAddress(repo) {
			return true, "", errors.New(fmt.Sprintf("repository '%v' does not exist", repo))
		}
		u, err := giturls.Parse(repo)
		if err != nil {
			return true, "", err
		}
		f_dir := u.Host + u.Path
		if ExistFile(DEFAULT_DIR_NAME) {
			f_dir = path.Join(DEFAULT_DIR_NAME, f_dir)
		}
		var out string
		if ExistFile(path.Join(f_dir, ".git")) {
			out, err = execCommand("git", "pull", f_dir)
		} else {
			out, err = execCommand("git", "clone", repo, f_dir)
		}
		return true, out, err
	},
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

	stdout, stderr, err := PipeCommand(cmd)
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

func PipeCommand(cmd *exec.Cmd) (string, string, error) {
	var output bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout, cmd.Stderr = &output, &stderr

	if err := cmd.Start(); err != nil {
		return output.String(), stderr.String(), err
	}
	cmd.Wait()

	return output.String(), stderr.String(), nil
}

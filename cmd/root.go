package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"

	"github.com/urfave/cli"
	"github.com/whilp/git-urls"
)

var RootFlagAction = &FlagAction{
	Flag:     nil,
	FlagName: "",
	Action: func(c *cli.Context, f *FlagAction) (bool, error) {
		if len(c.Args()) == 0 {
			// print usage
			return true, errors.New("No args!") // error
		}
		repo := c.Args().Get(0)
		if !ValidGitAddress(repo) {
			return true, errors.New(fmt.Sprintf("repository '%v' does not exist", repo))
		}
		u, err := giturls.Parse(repo)
		if err != nil {
			return true, err
		}
		f_dir := u.Host + u.Path
		if ExistFile(DEFAULT_DIR_NAME) {
			f_dir = path.Join(DEFAULT_DIR_NAME, f_dir)
		}
		//_, f_path := path.Split(f_dir)
		//f_path = strings.Split(f_path, ".")[0]
		if len(f_dir) != 0 /* && len(f_path) != 0*/ && !ExistFile(f_dir) {
			if err := os.MkdirAll(f_dir, os.ModePerm); err != nil {
				return true, err
			}
		}

		return true, nil
	},
}

func ValidGitAddress(repo string) bool {
	matched, err := regexp.MatchString(
		`((git|ssh|http(s)?)|(git@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(\.git)(/)?`,
		repo)
	return matched && err == nil
}

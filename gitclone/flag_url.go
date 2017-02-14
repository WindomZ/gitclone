package gitclone

import (
	"github.com/WindomZ/cli-mate"
	"github.com/urfave/cli"
)

var UrlFlag = cli_mate.Flag{
	Flag: cli.StringFlag{
		Name:  "url, u, ",
		Usage: "git repository `URL`",
	},
	FlagName: "url",
	Action: func(c *cli_mate.Context, f *cli_mate.Flag) cli_mate.ExitCoder {
		value := c.String(f.Name())
		if len(value) == 0 {
			return nil
		}
		var repo string
		if c.NArg() > 0 {
			repo = c.Args().Get(0)
		}
		content, err := urlAction(repo)
		if err != nil {
			return cli_mate.NewExitError(err)
		}
		c.Println(content)
		return nil
	},
}

func urlAction(repo string) (string, error) {
	return rootAction(repo)
}

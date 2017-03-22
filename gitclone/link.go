package gitclone

import (
	"github.com/WindomZ/go-commander"
)

var LinkAction = func(c commander.Context) error {
	filepath := c.GetString("<filepath>")
	println(filepath)
	return nil
}

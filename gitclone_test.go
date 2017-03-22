package main

import (
	"github.com/WindomZ/testify/assert"
	"os"
	"os/exec"
	"path"
	"testing"
)

func TestGitClone_Root(t *testing.T) {
	cmd := exec.Command("gitclone", "https://github.com/WindomZ/gitclone.git")
	if err := cmd.Run(); err != nil {
		assert.Error(t, err)
	}

	_, err := os.Stat(path.Join("github.com", "WindomZ", "gitclone", ".git"))
	assert.Equal(t,
		err == nil || os.IsExist(err),
		true,
		"No file exists: './GitClones/github.com/WindomZ/gitclone/.git'",
	)
}

func TestGitClone_Clean(t *testing.T) {
	cmd := exec.Command("rm", "-rf", "github.com")
	if err := cmd.Run(); err != nil {
		assert.Error(t, err)
	}
}

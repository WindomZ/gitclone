package main

import (
	"github.com/WindomZ/gitclone/gitclone"
	"github.com/WindomZ/testify/assert"
	"os/exec"
	"path"
	"testing"
)

func TestGitClone_Root(t *testing.T) {
	cmd := exec.Command("gitclone", "https://github.com/WindomZ/gitclone.git")
	if err := cmd.Run(); err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t,
		gitclone.ExistFile(path.Join("github.com", "WindomZ", "gitclone", ".git")),
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

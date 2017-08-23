package main

import (
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/WindomZ/testify/assert"
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

func TestGitClone_Root_Depth(t *testing.T) {
	cmd := exec.Command("gitclone", "https://github.com/WindomZ/gitclone.git", "--depth", "2")
	if err := cmd.Run(); err != nil {
		assert.Error(t, err)
	}

	_, err := os.Stat(path.Join("WindomZ", "gitclone", ".git"))
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

	cmd = exec.Command("rm", "-rf", "WindomZ")
	if err := cmd.Run(); err != nil {
		assert.Error(t, err)
	}
}

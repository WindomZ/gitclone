package tests

import (
	"github.com/WindomZ/gitclone/gitclone"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"path"
	"testing"
)

func TestGitClone_Mkdir(t *testing.T) {
	cmd := exec.Command("gitclone", "mkdir")
	if err := cmd.Run(); err != nil {
		assert.Error(t, err)
	}
	assert.Equal(
		t,
		gitclone.ExistFile(gitclone.DEFAULT_DIR_NAME),
		true,
		"Fail to create a default directory",
	)
}

func TestGitClone_Root(t *testing.T) {
	cmd := exec.Command("gitclone", "https://github.com/WindomZ/gitclone.git")
	if err := cmd.Run(); err != nil {
		assert.Error(t, err)
	}
	assert.Equal(
		t,
		gitclone.ExistFile(path.Join(gitclone.DEFAULT_DIR_NAME, "github.com", "WindomZ", "gitclone", ".git")),
		true,
		"No file exists: './GitClones/github.com/WindomZ/gitclone/.git'",
	)
}

func TestGitClone_Clean(t *testing.T) {
	cmd := exec.Command("rm", "-rf", gitclone.DEFAULT_DIR_NAME)
	if err := cmd.Run(); err != nil {
		assert.Error(t, err)
	}
}

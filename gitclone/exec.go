package gitclone

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

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

func execChdir(dir string) error {
	return os.Chdir(dir)
}

package checks

import (
	"os"
	"os/exec"
	"strings"
)

func Command(command string) (string, error) {

	cmd := exec.Command(
		"/bin/bash",
		"-c",
		command,
	)

	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()

	return strings.TrimSpace(string(out)), err
}

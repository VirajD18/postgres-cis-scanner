package controltemplatepath

import (
	"os"
	"path/filepath"
)

func Resolve(path string) string {

	// If the configured path already exists, use it.
	if _, err := os.Stat(path); err == nil {
		return path
	}

	// RPM installation path.
	base := filepath.Base(path)

	installed := filepath.Join("/etc/pgcis/templates", base)
	if _, err := os.Stat(installed); err == nil {
		return installed
	}

	// Development path relative to executable.
	if exe, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exe)

		local := filepath.Join(exeDir, path)
		if _, err := os.Stat(local); err == nil {
			return local
		}
	}

	return path
}

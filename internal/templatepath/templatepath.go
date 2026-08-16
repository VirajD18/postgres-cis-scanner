package templatepath

import (
	"os"
	"path/filepath"
)

func Resolve(name string) string {

	// Optional override.
	if dir := os.Getenv("PGCIS_TEMPLATE_DIR"); dir != "" {
		path := filepath.Join(dir, name)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// RPM installation path.
	installed := filepath.Join("/usr/share/pgcis/templates", name)
	if _, err := os.Stat(installed); err == nil {
		return installed
	}

	// Development path relative to the executable.
	if exe, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exe)

		// Binary in project root:
		// /home/viraj/postgres-cis-scanner/pgcis
		// -> /home/viraj/postgres-cis-scanner/templates
		local := filepath.Join(exeDir, "templates", name)

		if _, err := os.Stat(local); err == nil {
			return local
		}
	}

	// Fallback for running with go run or from project directory.
	local := filepath.Join("templates", name)

	return local
}

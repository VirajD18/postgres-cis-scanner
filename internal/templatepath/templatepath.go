package templatepath

import (
	"os"
	"path/filepath"
)

func Resolve(name string) string {
	// Optional override for development/custom installations.
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

	// Development/source-tree path.
	local := filepath.Join("templates", name)
	if _, err := os.Stat(local); err == nil {
		return local
	}

	// Return the local path so the caller gets the normal
	// "file not found" error if none of the locations exist.
	return local
}

package benchmarkpath

import (
	"os"
	"path/filepath"
)

func Resolve(name string) string {

	// Optional override.
	if dir := os.Getenv("PGCIS_BENCHMARK_DIR"); dir != "" {
		path := filepath.Join(dir, name)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// RPM installation path.
	installed := filepath.Join("/usr/share/pgcis/benchmark", name)
	if _, err := os.Stat(installed); err == nil {
		return installed
	}

	// Development path relative to executable.
	if exe, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exe)

		local := filepath.Join(exeDir, "benchmark", name)

		if _, err := os.Stat(local); err == nil {
			return local
		}
	}

	// Fallback for go run or running from project directory.
	local := filepath.Join("benchmark", name)

	return local
}

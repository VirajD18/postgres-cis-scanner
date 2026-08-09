package inventory

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
)

func detectPlatform(conn *pgx.Conn) (string, bool, error) {

	var version string

	err := conn.QueryRow(
		context.Background(),
		"SELECT version()",
	).Scan(&version)

	if err != nil {
		return "", false, err
	}

	v := strings.ToLower(version)

	switch {

	case strings.Contains(v, "aurora"):
		return "aurora", true, nil

	case strings.Contains(v, "amazon rds"):
		return "rds", true, nil

	case strings.Contains(v, "azure"):
		return "azure-flex", true, nil

	default:
		return "self-managed", false, nil
	}
}

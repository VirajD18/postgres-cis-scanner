package inventory

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func detectVersion(conn *pgx.Conn) (string, error) {

	var version string

	err := conn.QueryRow(
		context.Background(),
		"SHOW server_version",
	).Scan(&version)

	if err != nil {
		return "", err
	}

	return version, nil
}

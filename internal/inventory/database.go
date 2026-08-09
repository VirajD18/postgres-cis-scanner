package inventory

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func detectDatabase(conn *pgx.Conn) (string, error) {

	var database string

	err := conn.QueryRow(
		context.Background(),
		"SELECT current_database()",
	).Scan(&database)

	if err != nil {
		return "", err
	}

	return database, nil
}

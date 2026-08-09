package checks

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Extension(conn *pgx.Conn, extension string) (bool, error) {

	var exists bool

	err := conn.QueryRow(
		context.Background(),
		`SELECT EXISTS (
			SELECT 1
			FROM pg_extension
			WHERE extname = $1
		);`,
		extension,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

package inventory

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func detectExtensions(conn *pgx.Conn) ([]string, error) {

	rows, err := conn.Query(
		context.Background(),
		"SELECT extname FROM pg_extension ORDER BY extname",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var extensions []string

	for rows.Next() {
		var ext string

		if err := rows.Scan(&ext); err != nil {
			return nil, err
		}

		extensions = append(extensions, ext)
	}

	return extensions, nil
}

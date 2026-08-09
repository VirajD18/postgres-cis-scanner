package checks

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Parameter(conn *pgx.Conn, parameter string) (string, error) {

	query := fmt.Sprintf("SHOW %s;", parameter)

	var value string

	err := conn.QueryRow(
		context.Background(),
		query,
	).Scan(&value)

	if err != nil {
		return "", err
	}

	return value, nil
}

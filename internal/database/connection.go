package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/VirajD18/postgres-cis-scanner/internal/models"
)

func Connect(cfg *models.Config) (*pgx.Conn, error) {

	connString := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.User,
		cfg.Password,
		cfg.SSLMode,
	)

	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

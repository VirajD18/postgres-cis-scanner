package inventory

import (
	"github.com/jackc/pgx/v5"

	"github.com/VirajD18/postgres-cis-scanner/internal/models"
)

func Build(conn *pgx.Conn, cfg *models.Config, server models.Server) (*models.Inventory, error) {

	inv := &models.Inventory{}

	inv.Host = cfg.Host

	inv.HAHosts = server.HAHosts
	inv.DRHost = server.DRHost

	version, err := detectVersion(conn)
	if err != nil {
		return nil, err
	}

	inv.PostgresVersion = version

	database, err := detectDatabase(conn)
	if err != nil {
		return nil, err
	}

	inv.DatabaseName = database

	platform, managed, err := detectPlatform(conn)
	if err != nil {
		return nil, err
	}

	// Respect explicitly configured platform type.
	switch server.Type {
	case "iaas":
		inv.Platform = "self-managed"
		inv.ManagedService = false

	case "rds":
		inv.Platform = "rds"
		inv.ManagedService = true

	case "aurora":
		inv.Platform = "aurora"
		inv.ManagedService = true

	case "azure-flex":
		inv.Platform = "azure-flex"
		inv.ManagedService = true

	default:
		// Fall back to database-side detection.
		inv.Platform = platform
		inv.ManagedService = managed
	}

	extensions, err := detectExtensions(conn)
	if err != nil {
		return nil, err
	}

	inv.Extensions = extensions
	inv.Modules = detectModules(inv.Platform)

	return inv, nil
}

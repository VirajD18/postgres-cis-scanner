package scanner

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/VirajD18/postgres-cis-scanner/internal/benchmark"
	"github.com/VirajD18/postgres-cis-scanner/internal/benchmarkpath"
	"github.com/VirajD18/postgres-cis-scanner/internal/charts"
	"github.com/VirajD18/postgres-cis-scanner/internal/dashboard"
	"github.com/VirajD18/postgres-cis-scanner/internal/database"
	"github.com/VirajD18/postgres-cis-scanner/internal/engine"
	"github.com/VirajD18/postgres-cis-scanner/internal/inventory"
	"github.com/VirajD18/postgres-cis-scanner/internal/models"
	"github.com/VirajD18/postgres-cis-scanner/internal/platform"
	"github.com/VirajD18/postgres-cis-scanner/internal/report"
	"github.com/VirajD18/postgres-cis-scanner/internal/servers"
)

func Scan(server models.Server) (dashboard.ServerReport, error) {

	cfg := &models.Config{
		Host:     server.Host,
		Port:     server.Port,
		Database: server.Database,
		User:     server.User,
		Password: server.Password,
		SSLMode:  server.SSLMode,
	}

	conn, err := database.Connect(cfg)
	if err != nil {
		return dashboard.ServerReport{}, err
	}
	defer conn.Close(context.Background())

	inv, err := inventory.Build(conn, cfg, server)
	if err != nil {
		return dashboard.ServerReport{}, err
	}

	// Determine PostgreSQL major version
	major := strings.Split(inv.PostgresVersion, ".")[0]

	benchmarkDir := benchmarkpath.Resolve(
		"PostgreSQL" + major,
	)

	// Fallback if benchmark directory doesn't exist
	if _, err := os.Stat(benchmarkDir); os.IsNotExist(err) {
		benchmarkDir = benchmarkpath.Resolve("PostgreSQL18")
	}

	// Store benchmark used
	inv.Benchmark = filepath.Base(benchmarkDir)

	controls, err := benchmark.Load(benchmarkDir)
	if err != nil {
		return dashboard.ServerReport{}, err
	}

	// Apply server-specific control template for IaaS servers
	if server.ControlTemplate != "" {

		allowed, err := servers.LoadControlTemplate(server.ControlTemplate)
		if err != nil {
			return dashboard.ServerReport{}, err
		}

		var filtered []models.Control

		for _, control := range controls {
			if allowed[control.ID] {
				filtered = append(filtered, control)
			}
		}

		controls = filtered
	}

	var results []models.Result

	for _, control := range controls {

		// Skip controls that don't apply to this platform
		if !platform.Supported(control, inv) {

			results = append(results, models.Result{
				ControlID: control.ID,
				Title:     control.Title,
				Severity:  control.Severity,
				Status:    "NOT_APPLICABLE",
				Message:   "Control not applicable for this platform",
			})

			continue
		}

		r, err := engine.Execute(conn, control)
		if err != nil {
			continue
		}

		results = append(results, *r)
	}

	safeName := strings.ReplaceAll(server.Name, " ", "_")

	outDir := filepath.Join("reports", safeName)

	if err := os.MkdirAll(outDir, 0755); err != nil {
		return dashboard.ServerReport{}, err
	}

	if err := report.SaveJSON(
		results,
		filepath.Join(outDir, "results.json"),
	); err != nil {
		return dashboard.ServerReport{}, err
	}

	if err := report.SaveHTML(
		inv,
		results,
		filepath.Join(outDir, "report.html"),
	); err != nil {
		return dashboard.ServerReport{}, err
	}

	if err := charts.Generate(results, outDir); err != nil {
		return dashboard.ServerReport{}, err
	}

	summary := report.BuildSummary(results)

	serverReport := dashboard.Build(
		inv,
		summary,
		server.Name,
	)

	return serverReport, nil
}

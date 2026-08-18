package dashboard

import (
	"time"

	"github.com/VirajD18/postgres-cis-scanner/internal/models"
)

func DisplayPlatform(platform string) string {
	switch platform {
	case "self-managed":
		return "Self-Managed"
	case "rds":
		return "RDS"
	case "aurora":
		return "Aurora"
	case "azure-flex":
		return "Azure-Flexible"
	default:
		return platform
	}
}

func Build(
	inv *models.Inventory,
	summary models.Summary,
	serverName string,
) ServerReport {

	return ServerReport{
		Name:       serverName,
		Version:    inv.PostgresVersion,
		Platform:   DisplayPlatform(inv.Platform),
		Compliance: summary.Compliance,
		Pass:       summary.Pass,
		Fail:       summary.Fail,
		Manual:     summary.Manual,
		Info:       summary.Info,
		LastScan:   time.Now().Format("2006-01-02 15:04:05"),
		ReportPath: serverName + "/report.html",
	}
}

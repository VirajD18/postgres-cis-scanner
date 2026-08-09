package platform

import "github.com/VirajD18/postgres-cis-scanner/internal/models"

func Supported(control models.Control, inventory *models.Inventory) bool {

	// No platform restriction
	if len(control.Platforms) == 0 {
		return true
	}

	for _, platform := range control.Platforms {
		if platform == inventory.Platform {
			return true
		}
	}

	return false
}

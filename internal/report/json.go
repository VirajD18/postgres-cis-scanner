package report

import (
	"encoding/json"
	"os"

	"github.com/VirajD18/postgres-cis-scanner/internal/models"
)

func SaveJSON(results []models.Result, file string) error {

	data, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0644)
}

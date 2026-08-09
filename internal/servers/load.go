package servers

import (
	"encoding/json"
	"os"

	"github.com/VirajD18/postgres-cis-scanner/internal/models"
)

func Load(path string) ([]models.Server, error) {

	var servers []models.Server

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &servers)
	if err != nil {
		return nil, err
	}

	return servers, nil
}

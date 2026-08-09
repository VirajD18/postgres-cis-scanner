package servers

import (
	"encoding/json"
	"os"

	"github.com/VirajD18/postgres-cis-scanner/internal/models"
)

func LoadControlTemplate(path string) (map[string]bool, error) {

	var template models.ControlTemplate

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(file, &template); err != nil {
		return nil, err
	}

	allowed := make(map[string]bool)

	for _, controlID := range template.Controls {
		allowed[controlID] = true
	}

	return allowed, nil
}

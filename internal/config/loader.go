package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/VirajD18/postgres-cis-scanner/internal/models"
)

func LoadConfig(path string) (*models.Config, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %w", err)
	}
	defer file.Close()

	var cfg models.Config

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid config file: %w", err)
	}

	return &cfg, nil
}

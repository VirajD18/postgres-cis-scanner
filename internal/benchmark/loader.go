package benchmark

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/VirajD18/postgres-cis-scanner/internal/models"
)

func Load(directory string) ([]models.Control, error) {

	var controls []models.Control

	files, err := filepath.Glob(filepath.Join(directory, "*.json"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {

		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}

		var c []models.Control

		err = json.NewDecoder(f).Decode(&c)
		f.Close()

		if err != nil {
			return nil, err
		}

		controls = append(controls, c...)
	}

	return controls, nil
}

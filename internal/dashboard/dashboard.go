package dashboard

import (
	"html/template"
	"os"
	"path/filepath"
)

type ServerReport struct {
	Name         string
	Version      string
	Platform     string
	Compliance   float64
	Pass         int
	Fail         int
	Manual       int
	Info         int
	LastScan     string
	ReportPath   string
}

func Generate(reports []ServerReport) error {

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		return err
	}

	err = os.MkdirAll("reports", 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join("reports", "index.html"))
	if err != nil {
		return err
	}

	defer f.Close()

	return tmpl.Execute(f, reports)
}

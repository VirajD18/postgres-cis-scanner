package dashboard

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/VirajD18/postgres-cis-scanner/internal/templatepath"
)

type ServerReport struct {
	Name       string
	Version    string
	Platform   string
	Compliance float64
	Pass       int
	Fail       int
	Manual     int
	Info       int
	LastScan   string
	ReportPath string
}

func Generate(reports []ServerReport) error {

	tmpl, err := template.ParseFiles(templatepath.Resolve("index.html"))
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

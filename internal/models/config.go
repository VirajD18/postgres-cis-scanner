package models

type Config struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Database        string `json:"database"`
	User            string `json:"user"`
	Password        string `json:"password"`
	SSLMode         string `json:"sslmode"`
	Output          string `json:"output"`
	ReportDirectory string `json:"report_directory"`
}

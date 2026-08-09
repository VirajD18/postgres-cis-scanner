package models

type Server struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
	SSLMode  string `json:"sslmode"`

	HAHosts []string `json:"ha_hosts"`
	DRHost  string   `json:"dr_host"`

	Benchmark string `json:"benchmark"`
	Type string `json:"type"`
	ControlTemplate string `json:"control_template"`
}

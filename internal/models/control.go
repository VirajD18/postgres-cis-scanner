package models

type Control struct {
	ID          string   `json:"id"`
	Group       string   `json:"group"`
	Title        string   `json:"title"`

	CheckType    string   `json:"check_type"`

	Query         string   `json:"query,omitempty"`
	Parameter     string   `json:"parameter,omitempty"`
	Extension     string   `json:"extension,omitempty"`
	Command       string   `json:"command,omitempty"`

	Expected      string   `json:"expected"`
	Validation    string   `json:"validation"`

	Severity      string   `json:"severity"`

	Platforms     []string `json:"platforms"`

	Rationale     string   `json:"rationale"`
	Remediation   string   `json:"remediation"`
	Reference     string   `json:"reference"`

	Enabled       bool     `json:"enabled"`
}

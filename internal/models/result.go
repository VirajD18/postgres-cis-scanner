package models

type Result struct {
	ControlID string
	Group     string

	Title    string
	Severity string

	Status  string
	Message string

	Expected string
	Actual   string

	Rationale   string
	Remediation string
	Reference   string
}

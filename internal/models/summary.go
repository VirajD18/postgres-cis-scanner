package models

type Summary struct {
	Pass          int     `json:"pass"`
	Fail          int     `json:"fail"`
	Manual        int     `json:"manual"`
	Info          int     `json:"info"`
	NotApplicable int     `json:"not_applicable"`
	Error         int     `json:"error"`

	// Total contains only applicable controls: PASS + FAIL.
	Total      int     `json:"total"`
	Compliance float64 `json:"compliance"`
}

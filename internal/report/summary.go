package report

import "github.com/VirajD18/postgres-cis-scanner/internal/models"

func BuildSummary(results []models.Result) models.Summary {

	summary := models.Summary{}

	for _, result := range results {

		switch result.Status {

		case "PASS":
			summary.Pass++

		case "FAIL":
			summary.Fail++

		case "MANUAL":
			summary.Manual++

		case "INFO":
			summary.Info++

		case "NOT_APPLICABLE":
			summary.NotApplicable++

		case "ERROR":
			summary.Error++
		}
	}

	// Total represents only applicable controls.
	summary.Total = summary.Pass + summary.Fail

	// Compliance is calculated only from PASS and FAIL controls.
	applicable := summary.Pass + summary.Fail

	if applicable > 0 {
		summary.Compliance =
			(float64(summary.Pass) / float64(applicable)) * 100
	}

	return summary
}

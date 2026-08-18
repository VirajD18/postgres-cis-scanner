package report

import (
	"html/template"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/VirajD18/postgres-cis-scanner/internal/models"
	"github.com/VirajD18/postgres-cis-scanner/internal/templatepath"
)

type HTMLReport struct {
	Generated string
	Summary   models.Summary
	Inventory *models.Inventory

	Results       []HTMLResult
	FailedResults []HTMLResult
	GroupSummary  []GroupSummary
}

type HTMLResult struct {
	models.Result
	StatusClass string
}

type GroupSummary struct {
	Group  string
	Pass   int
	Fail   int
	Manual int
	Info   int
	NA     int
	Error  int
	Total  int
}

func SaveHTML(
	inventory *models.Inventory,
	results []models.Result,
	outputFile string,
) error {

	summary := BuildSummary(results)

	var htmlResults []HTMLResult
	var failed []HTMLResult

	groupMap := make(map[string]*GroupSummary)

	for _, r := range results {

		class := ""

		switch r.Status {

		case "PASS":
			class = "pass"

		case "FAIL":
			class = "fail"

		case "MANUAL":
			class = "manual"

		case "INFO":
			class = "info"

		case "NOT_APPLICABLE":
			class = "na"

		case "ERROR":
			class = "error"
		}

		hr := HTMLResult{
			Result:      r,
			StatusClass: class,
		}

		htmlResults = append(htmlResults, hr)

		if r.Status == "FAIL" {
			failed = append(failed, hr)
		}

		g, ok := groupMap[r.Group]
		if !ok {
			g = &GroupSummary{
				Group: r.Group,
			}
			groupMap[r.Group] = g
		}

		g.Total++

		switch r.Status {

		case "PASS":
			g.Pass++

		case "FAIL":
			g.Fail++

		case "MANUAL":
			g.Manual++

		case "INFO":
			g.Info++

		case "NOT_APPLICABLE":
			g.NA++

		case "ERROR":
			g.Error++
		}
	}

	var groups []GroupSummary

	for _, g := range groupMap {
		groups = append(groups, *g)
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Group < groups[j].Group
	})

	// Read organization logo and embed it directly into the HTML.
	//logoPath := "reports/Local_PostgreSQL/assets/logo.png"

	//logoData, err := os.ReadFile(logoPath)
	//if err != nil {
	//	return err
	//}

	//logoBase64 := base64.StdEncoding.EncodeToString(logoData)

	report := HTMLReport{
		Generated:     time.Now().Format("02-Jan-2006 15:04:05"),
		Summary:       summary,
		Inventory:     inventory,
		Results:       htmlResults,
		FailedResults: failed,
		GroupSummary:  groups,
	}

	funcMap := template.FuncMap{
		"lower": strings.ToLower,
		"displayPlatform": func(platform string) string {
			switch platform {
			case "self-managed":
				return "Self-Managed"
			case "rds":
				return "RDS"
			case "aurora":
				return "Aurora"
			case "azure-flex":
				return "Azure-Flexible"
			default:
				return platform
			}
		},
		"add": func(a, b int) int {
			return a + b
		},
		"add3": func(a, b, c int) int {
			return a + b + c
		},
	}

	tmpl, err := template.New("report.html").
		Funcs(funcMap).
		ParseFiles(templatepath.Resolve("report.html"))

	if err != nil {
		return err
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}

	defer f.Close()

	return tmpl.Execute(f, report)
}

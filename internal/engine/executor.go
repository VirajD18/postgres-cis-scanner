package engine

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/VirajD18/postgres-cis-scanner/internal/engine/checks"
	"github.com/VirajD18/postgres-cis-scanner/internal/engine/compare"
	"github.com/VirajD18/postgres-cis-scanner/internal/models"
)

func Execute(conn *pgx.Conn, control models.Control) (*models.Result, error) {

	result := &models.Result{
		ControlID:   control.ID,
		Group:       control.Group,
		Title:       control.Title,
		Severity:    control.Severity,
		Expected:    control.Expected,
		Rationale:   control.Rationale,
		Remediation: control.Remediation,
		Reference:   control.Reference,
	}

	switch control.CheckType {

	case "manual":

		result.Status = "MANUAL"
		result.Message = "Manual verification required"
		return result, nil

	case "linux":

	value, err := checks.Command(control.Command)

	if err != nil {
		result.Status = "ERROR"
		result.Message = err.Error()
		result.Actual = value
		return result, nil
	}

	result.Actual = value

	if compare.Evaluate(
		result.Actual,
		control.Expected,
		control.Validation,
	) {
		result.Status = "PASS"
	} else {
		result.Status = "FAIL"
	}

	return result, nil

	case "command":

		value, err := checks.Command(control.Command)

		if err != nil {
			result.Status = "ERROR"
			result.Message = err.Error()
			result.Actual = value
			return result, nil
		}

		result.Actual = value

		if compare.Evaluate(
			result.Actual,
			control.Expected,
			control.Validation,
		) {
			result.Status = "PASS"
		} else {
			result.Status = "FAIL"
		}

		return result, nil

	case "parameter":

		value, err := checks.Parameter(
			conn,
			control.Parameter,
		)

		if err != nil {
			result.Status = "ERROR"
			result.Message = err.Error()
			return result, nil
		}

		result.Actual = value

		if compare.Evaluate(
			result.Actual,
			control.Expected,
			control.Validation,
		) {
			result.Status = "PASS"
		} else {
			result.Status = "FAIL"
		}

		return result, nil

	case "extension":

		exists, err := checks.Extension(
			conn,
			control.Extension,
		)

		if err != nil {
			result.Status = "ERROR"
			result.Message = err.Error()
			return result, nil
		}

		if exists {
			result.Actual = "true"
		} else {
			result.Actual = "false"
		}

		if compare.Evaluate(
			result.Actual,
			control.Expected,
			control.Validation,
		) {
			result.Status = "PASS"
		} else {
			result.Status = "FAIL"
		}

		return result, nil

	case "sql":

		rows, err := conn.Query(
			context.Background(),
			control.Query,
		)

		if err != nil {
			result.Status = "ERROR"
			result.Message = err.Error()
			return result, nil
		}

		defer rows.Close()

		if !rows.Next() {
			result.Status = "NOT_APPLICABLE"
			return result, nil
		}

		fields := rows.FieldDescriptions()

		if len(fields) > 1 {
			result.Status = "MANUAL"
			result.Actual = "Review Required"
			return result, nil
		}

		var value interface{}

		if err := rows.Scan(&value); err != nil {
			result.Status = "ERROR"
			result.Message = err.Error()
			return result, nil
		}

		result.Actual = fmt.Sprintf("%v", value)

		if compare.Evaluate(
			result.Actual,
			control.Expected,
			control.Validation,
		) {
			result.Status = "PASS"
		} else {
			result.Status = "FAIL"
		}

		return result, nil
	}

	result.Status = "UNKNOWN"
	result.Message = "Unsupported check type"

	return result, nil
}

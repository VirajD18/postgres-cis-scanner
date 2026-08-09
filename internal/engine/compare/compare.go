package compare

import (
	"regexp"
	"strconv"
	"strings"
)

func Evaluate(actual, expected, validation string) bool {

	actual = strings.TrimSpace(actual)
	expected = strings.TrimSpace(expected)

	switch validation {

	case "equals":
		return strings.EqualFold(actual, expected)

	case "not_equals":
		return !strings.EqualFold(actual, expected)

	case "contains":
		return strings.Contains(
			strings.ToLower(actual),
			strings.ToLower(expected),
		)

	case "not_contains":
		return !strings.Contains(
			strings.ToLower(actual),
			strings.ToLower(expected),
		)

	case "starts_with":
		return strings.HasPrefix(
			strings.ToLower(actual),
			strings.ToLower(expected),
		)

	case "ends_with":
		return strings.HasSuffix(
			strings.ToLower(actual),
			strings.ToLower(expected),
		)

	case "empty":
		return actual == ""

	case "not_empty":
		return actual != ""

	case "in":

		values := strings.Split(expected, ",")

		for _, v := range values {
			if strings.EqualFold(
				actual,
				strings.TrimSpace(v),
			) {
				return true
			}
		}

		return false

	case "not_in":

		values := strings.Split(expected, ",")

		for _, v := range values {
			if strings.EqualFold(
				actual,
				strings.TrimSpace(v),
			) {
				return false
			}
		}

		return true

	case "greater_than":

		a, err1 := strconv.ParseFloat(actual, 64)
		e, err2 := strconv.ParseFloat(expected, 64)

		return err1 == nil &&
			err2 == nil &&
			a > e

	case "greater_or_equal":

		a, err1 := strconv.ParseFloat(actual, 64)
		e, err2 := strconv.ParseFloat(expected, 64)

		return err1 == nil &&
			err2 == nil &&
			a >= e

	case "less_than":

		a, err1 := strconv.ParseFloat(actual, 64)
		e, err2 := strconv.ParseFloat(expected, 64)

		return err1 == nil &&
			err2 == nil &&
			a < e

	case "less_or_equal":

		a, err1 := strconv.ParseFloat(actual, 64)
		e, err2 := strconv.ParseFloat(expected, 64)

		return err1 == nil &&
			err2 == nil &&
			a <= e

	case "regex":

		ok, err := regexp.MatchString(expected, actual)

		return err == nil && ok

	default:
		return false
	}
}

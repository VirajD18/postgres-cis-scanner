package inventory

import "strings"

func contextContains(s, sub string) bool {
	return strings.Contains(
		strings.ToLower(s),
		strings.ToLower(sub),
	)
}

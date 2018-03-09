package linkedin

import (
	"strings"
)

// Function removes all whitespaces, tabulators and newlines from endpoint's URL.
// Combines Endpoint URL with its parameters.
func PrepareURL(endpoint string, params string) string {
	return strings.TrimSpace(endpoint) + strings.Join(strings.Fields(params), "") + "?format=json"
}

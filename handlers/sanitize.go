package handlers

import "regexp"

var (
	nullArgPattern   = regexp.MustCompile(`\w+\s*:\s*\bnull\b`)
	extraCommaStart  = regexp.MustCompile(`\(\s*,\s*`)
	extraCommaEnd    = regexp.MustCompile(`\s*,\s*\)`)
	extraCommaMiddle = regexp.MustCompile(`\s*,\s*,\s*`)
	emptyParens      = regexp.MustCompile(`\(\s*\)`)
)

// sanitizeNullArgs removes inline null arguments from a GraphQL query string.
// graphql-go v0.8.1 does not support inline null literals.
func sanitizeNullArgs(query string) string {
	result := nullArgPattern.ReplaceAllString(query, "")
	result = extraCommaMiddle.ReplaceAllString(result, ", ")
	result = extraCommaStart.ReplaceAllString(result, "(")
	result = extraCommaEnd.ReplaceAllString(result, ")")
	result = emptyParens.ReplaceAllString(result, "")
	return result
}

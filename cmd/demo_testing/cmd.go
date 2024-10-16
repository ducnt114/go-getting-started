package demo_testing

import "strings"

// ConcatWithPlus concatenates strings using the + operator
func ConcatWithPlus(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s
	}
	return result
}

// ConcatWithBuilder concatenates strings using strings.Builder
func ConcatWithBuilder(strs []string) string {
	var builder strings.Builder
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

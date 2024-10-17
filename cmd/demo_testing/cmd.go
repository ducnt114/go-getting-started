package demo_testing

import (
	"encoding/json"
	gjson "github.com/goccy/go-json"
	"strings"
)

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

func JsonUnmarshall(input []byte, out interface{}) {
	_ = json.Unmarshal(input, &out)
}

func JsonUnmarshallWithLib(input []byte, out interface{}) {
	_ = gjson.Unmarshal(input, &out)
}

package util

import (
	"fmt"
	"strings"
)

func GenerateQuerySQL(input []string, start int) string {
	result := []string{}
	for idx := range input {
		result = append(result, fmt.Sprintf("$%d", start+idx))
	}

	return strings.Join(result, ",")
}

package utils

import (
	"strconv"
	"strings"
)

func RemoveComments(s string) string {
	processed, _, _ := strings.Cut(s, "//")

	return strings.Trim(processed, " ")
}

func IsNumeric(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}

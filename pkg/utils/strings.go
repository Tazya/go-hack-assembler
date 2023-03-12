package utils

import "strings"

func RemoveComments(s string) string {
	processed, _, _ := strings.Cut(s, "//")

	return strings.Trim(processed, " ")
}

package solutions

import "strings"

func CheckString(s string) bool {
	return strings.Count(s, "ba") == 0
}

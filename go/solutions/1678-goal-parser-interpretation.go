package solutions

import "strings"

func Interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	return strings.ReplaceAll(command, "(al)", "al")
}

package solutions

import "strings"

func NumUniqueEmails(emails []string) int {
	set := make(map[string]bool)
	for _, email := range emails {
		parts := strings.Split(email, "@")
		first := strings.Split(parts[0], "+")
		final := strings.Replace(first[0], ".", "", -1) + "@" + parts[1]

		set[final] = true
	}

	return len(set)
}

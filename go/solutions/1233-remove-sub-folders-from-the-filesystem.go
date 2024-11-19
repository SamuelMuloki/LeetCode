package solutions

import (
	"sort"
	"strings"
)

func RemoveSubfolders(folder []string) []string {
	sort.Strings(folder)

	result := []string{}
	result = append(result, folder[0])
	for i := 1; i < len(folder); i++ {
		lastFolder := result[len(result)-1]
		lastFolder += string('/')

		if !strings.HasPrefix(folder[i], lastFolder) {
			result = append(result, folder[i])
		}
	}

	return result
}

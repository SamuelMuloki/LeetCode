package solutions

import "strings"

func SimplifyPath(path string) string {
	st := []string{}
	folders := strings.Split(path, "/")
	for i := range folders {
		if folders[i] == "." || folders[i] == "" {
			continue
		} else if folders[i] == ".." {
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
		} else {
			st = append(st, folders[i])
		}
	}

	return "/" + strings.Join(st, "/")
}

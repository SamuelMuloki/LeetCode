package solutions

import "strings"

func MinRemoveToMakeValid(s string) string {
	st := []int{}
	sb := []byte(s)
	for i := range sb {
		if sb[i] == '(' {
			st = append(st, i)
		} else if sb[i] == ')' {
			if len(st) > 0 {
				st = st[:len(st)-1]
			} else {
				sb[i] = '*'
			}
		}
	}

	for len(st) > 0 {
		sb[st[len(st)-1]] = '*'
		st = st[:len(st)-1]
	}

	return strings.ReplaceAll(string(sb), "*", "")
}

package solutions

import "strconv"

func LongestCommonPrefix2(arr1 []int, arr2 []int) int {
	prefix := make(map[string]bool)
	for _, val := range arr1 {
		str := strconv.Itoa(val)
		pref := ""
		for _, ch := range str {
			pref += string(ch)
			prefix[pref] = true
		}
	}

	ans := 0
	for _, val := range arr2 {
		str := strconv.Itoa(val)
		pref := ""
		for _, ch := range str {
			pref += string(ch)
			if _, ok := prefix[pref]; ok {
				ans = max(ans, len(pref))
			}
		}
	}

	return ans
}

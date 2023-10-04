package solutions

func StrStr(haystack string, needle string) int {
	nLen := len(needle)
	for i := 0; i < len(haystack); i++ {
		rem := len(haystack)
		if i+nLen < rem {
			rem = i + nLen
		}

		if haystack[i] == needle[0] && string(haystack[i:rem]) == needle {
			return i
		}
	}

	return -1
}

package solutions

func CanConstruct(s string, k int) bool {
	if k > len(s) {
		return false
	}

	count := [26]int{}
	for _, ch := range s {
		count[ch-'a']++
	}

	for _, num := range count {
		if num%2 == 0 {
			continue
		}
		k--
	}

	return k >= 0
}

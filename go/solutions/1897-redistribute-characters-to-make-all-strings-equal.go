package solutions

func MakeEqual(words []string) bool {
	set := [26]int{}
	for i := range words {
		for j := range words[i] {
			set[words[i][j]-'a']++
		}
	}

	for _, wLen := range set {
		if wLen%len(words) != 0 {
			return false
		}
	}

	return true
}

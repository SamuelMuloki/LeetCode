package solutions

func RepeatedCharacter(s string) byte {
	set := make(map[byte]int)
	for i := range s {
		set[s[i]]++

		if set[s[i]] > 1 {
			return s[i]
		}
	}

	return s[0]
}

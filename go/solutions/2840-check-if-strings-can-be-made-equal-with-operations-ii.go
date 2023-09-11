package solutions

func CheckStrings(s1 string, s2 string) bool {
	odd := make([]int, 26)
	even := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		if i&1 == 0 {
			odd[s1[i]-'a']++
			odd[s2[i]-'a']--
		} else {
			even[s1[i]-'a']++
			even[s2[i]-'a']--
		}
	}

	for k := 0; k < len(odd); k++ {
		if odd[k] != 0 || even[k] != 0 {
			return false
		}
	}

	return true
}

package solutions

func CompareVersion(version1 string, version2 string) int {
	r1, r2 := []rune(version1), []rune(version2)
	for i, j := 0, 0; i < len(r1) || j < len(r2); {
		var n1, n2 int
		for ; i < len(r1) && r1[i] != '.'; i++ {
			n1 = n1*10 + int(r1[i]-'0')
		}
		for ; j < len(r2) && r2[j] != '.'; j++ {
			n2 = n2*10 + int(r2[j]-'0')
		}
		if n1 > n2 {
			return 1
		} else if n1 < n2 {
			return -1
		}
		i++
		j++
	}

	return 0
}

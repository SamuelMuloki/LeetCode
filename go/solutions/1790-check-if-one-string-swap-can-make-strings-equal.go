package solutions

func AreAlmostEqual(s1 string, s2 string) bool {
	firstIndexOff, secondIndexOf := 0, 0
	numDiffs := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			numDiffs++

			if numDiffs > 2 {
				return false
			} else if numDiffs == 1 {
				firstIndexOff = i
			} else {
				secondIndexOf = i
			}
		}
	}

	return s1[firstIndexOff] == s2[secondIndexOf] &&
		s1[secondIndexOf] == s2[firstIndexOff]
}

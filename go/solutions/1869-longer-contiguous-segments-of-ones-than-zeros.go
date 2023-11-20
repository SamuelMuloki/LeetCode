package solutions

func CheckZeroOnes(s string) bool {
	maxZeroes, maxOnes := 0, 0
	zeroCnt, oneCnt := 0, 0
	for i := range s {
		if s[i] == '0' {
			zeroCnt++
			oneCnt = 0
		} else {
			oneCnt++
			zeroCnt = 0
		}

		maxZeroes = max(maxZeroes, zeroCnt)
		maxOnes = max(maxOnes, oneCnt)
	}

	return maxOnes > maxZeroes
}

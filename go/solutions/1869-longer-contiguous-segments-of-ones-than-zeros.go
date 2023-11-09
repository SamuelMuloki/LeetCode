package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

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

		maxZeroes = utils.Max(maxZeroes, zeroCnt)
		maxOnes = utils.Max(maxOnes, oneCnt)
	}

	return maxOnes > maxZeroes
}

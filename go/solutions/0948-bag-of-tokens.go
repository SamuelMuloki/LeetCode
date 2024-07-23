package solutions

import "sort"

func BagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	ans, currPower := 0, power

	i, j := 0, len(tokens)-1
	for i <= j {
		if currPower >= tokens[i] {
			currPower -= tokens[i]
			ans++
			i++
		} else if i < j && ans > 0 {
			currPower += tokens[j]
			ans--
			j--
		} else {
			return ans
		}
	}

	return ans
}

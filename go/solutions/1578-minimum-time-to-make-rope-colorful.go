package solutions

func MinCost(colors string, neededTime []int) int {
	curSum, curMax, res := 0, 0, 0
	for i := 0; i < len(colors); i++ {
		if i > 0 && colors[i] != colors[i-1] {
			res += curSum - curMax
			curSum, curMax = 0, 0
		}
		curSum += neededTime[i]
		curMax = max(curMax, neededTime[i])
	}
	return res + curSum - curMax
}

package solutions

import "math"

func CheckArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		ans[i] = checkArithmeticSub(nums[l[i] : r[i]+1])
	}

	return ans
}

func checkArithmeticSub(nums []int) bool {
	minELement, maxELement := math.MaxInt, math.MinInt
	set := make(map[int]int)

	for i := range nums {
		minELement = min(minELement, nums[i])
		maxELement = max(maxELement, nums[i])
		set[nums[i]]++
	}

	if (maxELement-minELement)%(len(nums)-1) != 0 {
		return false
	}

	diff := (maxELement - minELement) / (len(nums) - 1)
	curr := minELement + diff
	for curr < maxELement {
		if _, ok := set[curr]; !ok {
			return false
		}

		curr += diff
	}

	return true
}

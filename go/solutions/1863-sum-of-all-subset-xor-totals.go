package solutions

func SubsetXORSum(nums []int) int {
	ans := 0
	curr := make([]int, 0)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		currXor := 0
		for i := range curr {
			currXor ^= curr[i]
		}
		ans += currXor
		if idx == len(nums) {
			return
		}

		for i := idx; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0)

	return ans
}

package solutions

func Rob(nums []int) int {
	prev1, prev2 := 0, 0
	for i := range nums {
		temp := prev1
		prev1 = max(prev2+nums[i], prev1)
		prev2 = temp
	}

	return prev1
}

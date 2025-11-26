package solutions

func MinimumOperations4(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%3 == 0 {
			continue
		}
		res++
	}

	return res
}

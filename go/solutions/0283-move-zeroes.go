package solutions

func MoveZeroes(nums []int) {
	for lastNonZeroFoundAt, curr := 0, 0; curr < len(nums); curr++ {
		if nums[curr] != 0 {
			swapNums(nums, lastNonZeroFoundAt, curr)
			lastNonZeroFoundAt++
		}
	}
}

func swapNums(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

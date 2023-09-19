package solutions

func NumberOfPoints(nums [][]int) int {
	set := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		for j := nums[i][0]; j <= nums[i][1]; j++ {
			set[j] = true
		}
	}

	return len(set)
}

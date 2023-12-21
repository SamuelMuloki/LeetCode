package solutions

func RepeatedNTimes(nums []int) int {
	set := make(map[int]int)
	for i := range nums {
		set[nums[i]]++
		if set[nums[i]] > 1 {
			return nums[i]
		}
	}

	return -1
}

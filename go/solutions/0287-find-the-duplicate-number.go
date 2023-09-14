package solutions

func FindDuplicate(nums []int) int {
	set := make(map[int]int)

	i := 0
	for {
		set[nums[i]]++

		if set[nums[i]] > 1 {
			return nums[i]
		}

		i++
	}
}

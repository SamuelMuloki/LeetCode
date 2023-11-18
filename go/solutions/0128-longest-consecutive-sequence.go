package solutions

func LongestConsecutive(nums []int) int {
	set := make(map[int]bool, 0)
	for i := range nums {
		set[nums[i]] = true
	}

	count := 0
	for num := range set {
		if set[num-1] {
			continue
		}

		j, maxCount := num, 1
		for set[j+1] {
			maxCount++
			j++
		}
		count = max(count, maxCount)
	}

	return count
}

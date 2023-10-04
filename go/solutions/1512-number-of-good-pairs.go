package solutions

func NumIdenticalPairs(nums []int) int {
	set := make(map[int]int)
	for i := range nums {
		set[nums[i]]++
	}

	count := 0
	for _, val := range set {
		count += (val * (val - 1)) / 2
	}

	return count
}

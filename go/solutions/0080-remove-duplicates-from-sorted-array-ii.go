package solutions

func RemoveDuplicates2(nums []int) int {
	set := make(map[int]int)
	set[nums[0]]++

	count := 0
	for i := 1; i < len(nums); i++ {
		set[nums[i]]++
		if nums[i] == nums[i-1] && set[nums[i]] > 2 {
			count++
		} else {
			nums[i-count] = nums[i]
		}
	}

	return len(nums) - count
}

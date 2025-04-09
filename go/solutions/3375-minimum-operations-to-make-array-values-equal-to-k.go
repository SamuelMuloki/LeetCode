package solutions

func MinOperations9(nums []int, k int) int {
	set := [101]bool{}
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			return -1
		}
		if !set[nums[i]] && nums[i] != k {
			res++
		}
		set[nums[i]] = true

	}

	return res
}

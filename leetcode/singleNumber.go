package leetcode

func SingleNumber(nums []int) int {
	hash := make(map[int]int)

	var minVal int
	for idx := range nums {
		hash[nums[idx]]++
	}

	for i, k := range hash {
		if k == 1 {
			minVal = i
		}
	}

	return minVal
}

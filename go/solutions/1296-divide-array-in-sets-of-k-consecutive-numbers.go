package solutions

func IsPossibleDivide(nums []int, k int) bool {
	N := len(nums)
	if N%k != 0 {
		return false
	}

	count := make(map[int]int)
	for i := range nums {
		count[nums[i]]++
	}

	for len(count) > 0 {
		curr := smallest(count)
		for i := 0; i < k; i++ {
			if count[curr+i] == 0 {
				return false
			}
			count[curr+i]--
			if count[curr+i] < 1 {
				delete(count, curr+i)
			}
		}
	}

	return true
}

package solutions

func MinEmptyOperations(nums []int) int {
	set := make(map[int]int)
	for i := range nums {
		set[nums[i]]++
	}

	res := 0
	for _, count := range set {
		for count > 1 {
			if count%3 == 0 {
				count -= 3
			} else {
				count -= 2
			}
			res++
		}

		if count == 1 {
			return -1
		}
	}

	return res
}

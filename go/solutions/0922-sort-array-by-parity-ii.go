package solutions

func SortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	res := make([]int, len(nums))
	for i := range nums {
		if nums[i]%2 == 0 {
			res[even] = nums[i]
			even += 2
		} else {
			res[odd] = nums[i]
			odd += 2
		}
	}

	return res
}

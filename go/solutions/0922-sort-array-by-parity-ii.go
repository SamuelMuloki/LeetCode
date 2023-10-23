package solutions

func SortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	for even < len(nums) && odd < len(nums) {
		if nums[even]%2 == 0 {
			even += 2
		} else if nums[odd]%2 != 0 {
			odd += 2
		} else {
			nums[even], nums[odd] = nums[odd], nums[even]
			even += 2
			odd += 2
		}
	}

	return nums
}

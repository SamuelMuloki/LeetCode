package solutions

func BuildArray2(nums []int) []int {
	n := len(nums)
	// turn the array into a=nb+r
	for j := range nums {
		/* retrieve correct value from potentially already processed element
		i.e. get r(nums[j]) from something maybe already in the form a = nb + r
		if it isnt already processed (doesnt have nb yet), that's ok b/c
		r < n, so r % n will return the same value */
		nums[j] = nums[j] + n*(nums[nums[j]]%n)
	}

	for i := range nums {
		nums[i] = nums[i] / n
	}

	return nums
}

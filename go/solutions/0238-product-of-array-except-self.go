package solutions

func ProductExceptSelf(nums []int) []int {
	prod := make([]int, len(nums))

	prod[0] = 1
	for k := 1; k < len(nums); k++ {
		prod[k] = prod[k-1] * nums[k-1]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		prod[i] *= right
		right *= nums[i]
	}

	return prod
}

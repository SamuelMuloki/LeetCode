package solutions

func ProductExceptSelf(nums []int) []int {
	prod := make([]int, len(nums))

	for i := range nums {
		prod[i] = 1
	}

	temp := 1
	for i := 0; i < len(nums); i++ {
		prod[i] = temp
		temp *= nums[i]
	}

	temp = 1
	for i := len(nums) - 1; i >= 0; i-- {
		prod[i] *= temp
		temp *= nums[i]
	}

	return prod
}

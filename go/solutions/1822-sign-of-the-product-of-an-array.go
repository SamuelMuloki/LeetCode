package solutions

func ArraySign(nums []int) int {
	prod := 1
	for i := range nums {
		prod *= signFunc(nums[i])
	}

	return prod
}

func signFunc(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}

	return 0
}

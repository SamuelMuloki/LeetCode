package solutions

func FindGCD(nums []int) int {
	x, y := 0, 1001
	for i := range nums {
		x = max(x, nums[i])
		y = min(y, nums[i])
	}

	for y > 0 {
		x, y = y, x%y
	}

	return x
}

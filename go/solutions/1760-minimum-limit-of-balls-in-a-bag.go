package solutions

func MinimumSize(nums []int, maxOperations int) int {
	var canDivide = func(nums []int, maxBalls int, maxOperations int) bool {
		ops := 0
		for _, n := range nums {
			ops += (n+maxBalls-1)/maxBalls - 1
			if ops > maxOperations {
				return false
			}
		}
		return true
	}

	left, right := 1, 0
	for _, n := range nums {
		if n > right {
			right = n
		}
	}
	res := right

	for left < right {
		mid := left + (right-left)/2
		if canDivide(nums, mid, maxOperations) {
			right = mid
			res = right
		} else {
			left = mid + 1
		}
	}
	return res
}

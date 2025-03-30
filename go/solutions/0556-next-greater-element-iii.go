package solutions

import (
	"math"
	"strconv"
)

func NextGreaterElementIII(n int) int {
	var reverse = func(nums []byte) {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	var nextPermutation = func(nums []byte) bool {
		i := len(nums) - 2
		for i > 0 && nums[i+1] <= nums[i] {
			i--
		}

		if i < 0 {
			return false
		}

		j := len(nums) - 1
		for j > 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		reverse(nums[i+1:])

		return true
	}

	digits := []byte(strconv.Itoa(n))
	if !nextPermutation(digits) {
		return -1
	}

	res, _ := strconv.Atoi(string(digits))
	if res > math.MaxInt32 || res <= n {
		return -1
	}

	return res
}

package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func Candy(ratings []int) int {
	n := len(ratings)

	left := make([]int, len(ratings))
	for idx := range left {
		left[idx] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	candies := left[n-1]
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			left[i] = utils.Max(left[i], left[i+1]+1)
		}
		candies += left[i]
	}

	return candies
}

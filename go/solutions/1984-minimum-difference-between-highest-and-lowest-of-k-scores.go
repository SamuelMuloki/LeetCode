package solutions

import (
	"math"
	"sort"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MinimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	l := 0
	r := k - 1
	ans := math.MaxInt32

	for r < len(nums) {
		ans = utils.Min(ans, nums[r]-nums[l])
		l++
		r++
	}

	return ans
}

package solutions

import (
	"math"
	"sort"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MinimumSum(nums []int) int {
	res := math.MaxInt
	n := len(nums)
	arr := make([]int, 0)
	for i := 0; i < n-1; i++ {
		if len(arr) > 0 && arr[len(arr)-1] < nums[i] {
			rem := nums[i+1:]
			idx := sort.Search(len(rem), func(i int) bool {
				return rem[i] < nums[i]
			})

			if idx > -1 && rem[idx] < nums[i] {
				res = utils.Min(res, arr[len(arr)-1]+nums[i]+rem[idx])
			}

		} else {
			arr = append(arr, nums[i])
		}
	}

	if res == math.MaxInt {
		return -1
	}

	return res
}

package solutions

import (
	"math"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MaximumTripletValue(nums []int) int64 {
	maxV := math.MinInt
	for i := 0; i < len(nums)-2; i++ {
		for j := i; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				maxV = utils.Max(maxV, nums[k]*(nums[i]-nums[j]))
			}
		}
	}

	return int64(maxV)
}

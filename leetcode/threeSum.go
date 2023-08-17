// https://leetcode.com/problems/3sum/
package leetcode

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		for k, j := i+1, len(nums)-1; k < j; {
			n := nums[i] + nums[k] + nums[j]
			if n > 0 {
				j--
			} else if n < 0 {
				k++
			} else {
				result = append(result, []int{nums[i], nums[k], nums[j]})
				for k < j && nums[k] == nums[k+1] {
					k++
				}

				for k < j && nums[j] == nums[j-1] {
					j--
				}
				k++
				j--
			}
		}
	}
	return result
}

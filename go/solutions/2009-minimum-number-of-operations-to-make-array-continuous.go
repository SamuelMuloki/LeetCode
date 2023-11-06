package solutions

import (
	"sort"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MinContinousOperations(nums []int) int {
	n := len(nums)
	ops := n

	newArr := make([]int, len(nums))
	copy(newArr, nums)
	sort.Ints(newArr)

	unique := make([]int, 0)
	unique = append(unique, newArr[0])
	for i := 1; i < len(newArr); i++ {
		if newArr[i] != newArr[i-1] {
			unique = append(unique, newArr[i])
		}
	}

	j := 0
	for i := 0; i < len(unique); i++ {
		for j < len(unique) && unique[j] < unique[i]+n {
			j++
		}

		count := j - i
		ops = utils.Min(ops, n-count)
	}

	return ops
}

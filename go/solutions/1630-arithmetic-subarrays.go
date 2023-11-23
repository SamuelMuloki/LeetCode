package solutions

import "sort"

func CheckArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		sub := append([]int{}, nums[l[i]:r[i]+1]...)
		sort.Ints(sub)
		ans[i] = true
		for j := 2; j < len(sub); j++ {
			if sub[j]-sub[j-1] != sub[1]-sub[0] {
				ans[i] = false
				break
			}
		}
	}

	return ans
}

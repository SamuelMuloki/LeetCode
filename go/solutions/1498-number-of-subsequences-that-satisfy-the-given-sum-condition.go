package solutions

import "sort"

func NumSubseq(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	res, l, r := 0, 0, n-1
	mod := int(1e9 + 7)

	pows := make([]int, n)
	pows[0] = 1
	for i := 1; i < n; i++ {
		pows[i] = pows[i-1] * 2 % mod
	}

	for l <= r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			res = (res + pows[r-l]) % mod
			l++
		}
	}

	return res
}

package solutions

import "strconv"

func FindDifferentBinaryString(nums []string) string {
	set := make(map[int]bool)
	for i := range nums {
		num, _ := strconv.ParseInt(nums[i], 2, 64)
		set[int(num)] = true
	}

	num, n := 0, len(nums)
	for i := 0; i <= n; i++ {
		if _, ok := set[i]; !ok {
			num = i
			break
		}
	}

	ans := strconv.FormatInt(int64(num), 2)
	for len(ans) < n {
		ans = "0" + ans
	}

	return ans
}

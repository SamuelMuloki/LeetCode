package solutions

func MaxProduct(nums []int) int {
	n := len(nums)
	l, r := 1, 1
	ans := nums[0]
	for i := 0; i < n; i++ {
		if l == 0 {
			l = 1
		}

		if r == 0 {
			r = 1
		}

		l *= nums[i]
		r *= nums[n-1-i]

		ans = max(ans, max(l, r))
	}

	return ans
}

package solutions

func MaximumSum(nums []int) int {
	count := make(map[int]int)
	ans := -1
	for i := 0; i < len(nums); i++ {
		curr, num := 0, nums[i]
		for num > 0 {
			curr += (num % 10)
			num /= 10
		}

		if val, ok := count[curr]; ok {
			ans = max(ans, val+nums[i])
		}

		count[curr] = max(count[curr], nums[i])
	}

	return ans
}

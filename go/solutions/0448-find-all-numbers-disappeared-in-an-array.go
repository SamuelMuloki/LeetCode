package solutions

func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	ans, set := make([]int, 0), make(map[int]bool)
	for i := range nums {
		set[nums[i]] = true
	}

	for j := 1; j <= n; j++ {
		if !set[j] {
			ans = append(ans, j)
		}
	}

	return ans
}

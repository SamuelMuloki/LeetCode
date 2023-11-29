package solutions

func TotalHammingDistance(nums []int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		ones := 0
		for _, num := range nums {
			ones += (num >> i) & 1
		}

		zeroes := len(nums) - ones
		ans += ones * zeroes
	}

	return ans
}

package solutions

func TupleSameProduct(nums []int) int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			count[nums[i]*nums[j]]++
		}
	}

	ans := 0
	for _, freq := range count {
		pairs := ((freq - 1) * freq) / 2
		ans += 8 * pairs
	}

	return ans
}

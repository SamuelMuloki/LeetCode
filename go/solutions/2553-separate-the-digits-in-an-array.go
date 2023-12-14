package solutions

func SeparateDigits(nums []int) []int {
	ans := make([]int, 0)
	for i := range nums {
		curr := make([]int, 0)
		for num := nums[i]; num > 0; num /= 10 {
			curr = append([]int{num % 10}, curr...)
		}
		ans = append(ans, curr...)
	}

	return ans
}

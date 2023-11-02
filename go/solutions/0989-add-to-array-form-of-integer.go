package solutions

func AddToArrayForm(num []int, k int) []int {
	sum := k
	ans := make([]int, 0)
	for i := len(num) - 1; i >= 0 || sum > 0; i-- {
		if i >= 0 {
			sum += num[i]
		}

		ans = append([]int{sum % 10}, ans...)
		sum /= 10
	}

	return ans
}

package solutions

func FindMatrix(nums []int) [][]int {
	freq := make([]int, len(nums)+1)
	ans := [][]int{}
	for _, num := range nums {
		if freq[num] >= len(ans) {
			ans = append(ans, []int{})
		}

		ans[freq[num]] = append(ans[freq[num]], num)
		freq[num]++
	}

	return ans
}

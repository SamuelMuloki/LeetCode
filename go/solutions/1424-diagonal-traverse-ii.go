package solutions

func FindDiagonalOrder(nums [][]int) []int {
	groups := make(map[int][]int)
	for row := len(nums) - 1; row >= 0; row-- {
		for col := 0; col < len(nums[row]); col++ {
			diagonal := row + col
			groups[diagonal] = append(groups[diagonal], nums[row][col])
		}
	}

	ans := make([]int, 0)
	curr := 0
	for len(groups[curr]) > 0 {
		ans = append(ans, groups[curr]...)
		curr++
	}

	return ans
}

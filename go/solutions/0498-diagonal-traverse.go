package solutions

func FindDiagonalOrder(mat [][]int) []int {
	groups := make(map[int][]int)
	for row := len(mat) - 1; row >= 0; row-- {
		for col := 0; col < len(mat[row]); col++ {
			diagonal := row + col
			groups[diagonal] = append(groups[diagonal], mat[row][col])
		}
	}

	ans := make([]int, 0)
	curr := 0
	for len(groups[curr]) > 0 {
		if curr%2 == 0 {
			ans = append(ans, groups[curr]...)
		} else {
			arr := groups[curr]
			for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
				arr[i], arr[j] = arr[j], arr[i]
			}
			ans = append(ans, arr...)
		}
		curr++
	}

	return ans
}

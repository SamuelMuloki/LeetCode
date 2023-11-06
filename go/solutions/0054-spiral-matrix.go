package solutions

func SpiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0)

	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	dr, dc := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	x, y, di := 0, 0, 0
	for i := 0; i < m*n; i++ {
		ans = append(ans, matrix[x][y])

		seen[x][y] = true
		newX, newY := x+dr[di], y+dc[di]

		if 0 <= newX && newX < m && 0 <= newY && newY < n && !seen[newX][newY] {
			x, y = newX, newY
		} else {
			di = (di + 1) % 4
			x += dr[di]
			y += dc[di]
		}
	}

	return ans
}

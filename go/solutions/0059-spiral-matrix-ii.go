package solutions

func GenerateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	dr, dc := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	x, y, di := 0, 0, 0
	cnt := 1
	for i := 0; i < n*n; i++ {
		ans[x][y] = cnt
		cnt++

		seen[x][y] = true
		newX, newY := x+dr[di], y+dc[di]

		if 0 <= newX && newX < n && 0 <= newY && newY < n && !seen[newX][newY] {
			x, y = newX, newY
		} else {
			di = (di + 1) % 4
			x += dr[di]
			y += dc[di]
		}
	}

	return ans
}

package solutions

func ImageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum, count := 0, 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if 0 <= x && x < m && 0 <= y && y < n {
						sum += img[x][y]
						count++
					}
				}
			}

			ans[i][j] = sum / count
		}
	}

	return ans
}

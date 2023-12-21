package solutions

func CountPoints(points [][]int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i := range queries {
		x, y, r := queries[i][0], queries[i][1], queries[i][2]
		for _, point := range points {
			if sq(point[0]-x)+sq(point[1]-y) <= sq(r) {
				ans[i]++
			}
		}
	}

	return ans
}

func sq(x int) int {
	return x * x
}

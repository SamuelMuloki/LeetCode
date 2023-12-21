package solutions

import "sort"

func MaxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	ans := 0
	for i := 1; i < len(points); i++ {
		ans = max(ans, points[i][0]-points[i-1][0])
	}

	return ans
}

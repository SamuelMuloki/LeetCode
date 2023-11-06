package solutions

import (
	"sort"
)

func FindMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	ans, arrowPos := 1, points[0][1]
	for i := 1; i < len(points); i++ {
		if arrowPos >= points[i][0] {
			continue
		}
		ans++
		arrowPos = points[i][1]
	}

	return ans
}

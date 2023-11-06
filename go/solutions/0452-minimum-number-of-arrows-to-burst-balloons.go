package solutions

import (
	"sort"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func FindMinArrowShots(points [][]int) int {
	overlaps := make([][]int, 0)
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	overlaps = append(overlaps, points[0])
	for i := 1; i < len(points); i++ {
		n := len(overlaps)
		if points[i][0] <= overlaps[n-1][1] {
			overlaps[n-1][1] = utils.Min(overlaps[n-1][1], points[i][1])
		} else {
			overlaps = append(overlaps, points[i])
		}
	}

	return len(overlaps)
}

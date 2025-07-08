package solutions

import "sort"

func CheckValidCuts(n int, rectangles [][]int) bool {
	var check = func(dim int) bool {
		gapCount := 0
		sort.Slice(rectangles, func(i, j int) bool {
			return rectangles[i][dim] < rectangles[j][dim]
		})

		furthestEnd := rectangles[0][dim+2]
		for i := 1; i < len(rectangles); i++ {
			rect := rectangles[i]

			if furthestEnd <= rect[dim] {
				gapCount++
			}

			furthestEnd = max(furthestEnd, rect[dim+2])
		}

		return gapCount >= 2
	}

	return check(0) || check(1)
}

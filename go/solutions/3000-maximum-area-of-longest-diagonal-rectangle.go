package solutions

import "math"

func AreaOfMaxDiagonal(dimensions [][]int) int {
	currMax := float64(0)
	res := 0
	for _, dim := range dimensions {
		diag := math.Sqrt(float64(dim[0]*dim[0] + dim[1]*dim[1]))
		if diag > currMax {
			res = dim[0] * dim[1]
			currMax = diag
		} else if diag == currMax {
			res = max(res, dim[0]*dim[1])
		}
	}

	return res
}

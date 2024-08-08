package solutions

func SpiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	traversed := make([][]int, 0)

	for step, direction := 1, 0; len(traversed) < rows*cols; {
		for i := 0; i < 2; i++ {
			for j := 0; j < step; j++ {
				if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
					traversed = append(traversed, []int{rStart, cStart})
				}

				rStart += dir[direction][0]
				cStart += dir[direction][1]
			}

			direction = (direction + 1) % 4
		}
		step++
	}

	return traversed
}

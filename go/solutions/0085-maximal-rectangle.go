package solutions

func MaximalRectangle(matrix [][]byte) int {
	heights := make([]int, len(matrix[0])+1)
	heights[len(heights)-1] = -1
	ans := 0
	for _, row := range matrix {
		for i := range row {
			if row[i] == '1' {
				heights[i]++
			} else {
				heights[i] = 0
			}
		}

		stack := []int{}
		for i, currentHeight := range heights {
			for len(stack) > 0 && heights[stack[len(stack)-1]] > currentHeight {
				prev := heights[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]

				width := i
				if len(stack) > 0 {
					width = i - stack[len(stack)-1] - 1
				}
				ans = max(ans, prev*width)
			}
			stack = append(stack, i)
		}
	}

	return ans
}

package solutions

func SnakesAndLadders(board [][]int) int {
	n := len(board)
	visited := make([]bool, n*n+1)
	queue := []int{1}
	visited[1] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			pos := queue[0]
			queue = queue[1:]

			if pos == n*n {
				return steps
			}

			for j := 1; j <= 6 && pos+j <= n*n; j++ {
				nextPos := pos + j
				row, col := getRC(nextPos, n)
				if board[row][col] != -1 {
					nextPos = board[row][col]
				}
				if !visited[nextPos] {
					visited[nextPos] = true
					queue = append(queue, nextPos)
				}
			}
		}
		steps++
	}

	return -1
}

func getRC(pos, n int) (int, int) {
	quot := (pos - 1) / n
	row := n - 1 - quot
	col := (pos - 1) % n
	if quot%2 == 1 {
		col = n - 1 - col
	}
	return row, col
}

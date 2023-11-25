package solutions

type DiagonalPair struct {
	row, col int
}

func FindDiagonalOrder2(nums [][]int) []int {
	queue := make([]DiagonalPair, 0)
	queue = append(queue, DiagonalPair{0, 0})
	ans := make([]int, 0)
	for len(queue) > 0 {
		row, col := queue[0].row, queue[0].col
		queue = queue[1:]
		ans = append(ans, nums[row][col])

		if col == 0 && row+1 < len(nums) {
			queue = append(queue, DiagonalPair{row: row + 1, col: col})
		}

		if col+1 < len(nums[row]) {
			queue = append(queue, DiagonalPair{row: row, col: col + 1})
		}
	}

	return ans
}

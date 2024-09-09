package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SpiralMatrix(m int, n int, head *utils.ListNode) [][]int {
	i, j, cur_d := 0, 0, 0
	movement := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = -1
		}
	}

	for head != nil {
		res[i][j] = head.Val
		newi := i + movement[cur_d][0]
		newj := j + movement[cur_d][1]

		if min(newi, newj) < 0 || newi >= m || newj >= n || res[newi][newj] != -1 {
			cur_d = (cur_d + 1) % 4
		}

		i += movement[cur_d][0]
		j += movement[cur_d][1]

		head = head.Next
	}

	return res
}

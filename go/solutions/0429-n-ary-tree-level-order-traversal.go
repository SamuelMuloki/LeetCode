package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func NArryLevelOrder(root *utils.Node) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	queue := make([]*utils.Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		qLen := len(queue)
		curr := make([]int, 0)
		for i := 0; i < qLen; i++ {
			curr = append(curr, queue[i].Val)
			queue = append(queue, queue[i].Children...)
		}

		ans = append(ans, append([]int{}, curr...))
		queue = queue[qLen:]
	}

	return ans
}

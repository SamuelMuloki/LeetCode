package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Level struct {
	parent *utils.TreeNode
	node   *utils.TreeNode
}

func ReplaceValueInTree(root *utils.TreeNode) *utils.TreeNode {
	queue := make([]Level, 0)
	queue = append(queue, Level{parent: nil, node: root})
	for len(queue) > 0 {
		qLen := len(queue)
		curr := make(map[*utils.TreeNode]int)
		sum := 0
		for i := 0; i < qLen; i++ {
			sum += queue[i].node.Val
			curr[queue[i].parent] += queue[i].node.Val

			if queue[i].node.Left != nil {
				queue = append(queue, Level{parent: queue[i].node, node: queue[i].node.Left})
			}

			if queue[i].node.Right != nil {
				queue = append(queue, Level{parent: queue[i].node, node: queue[i].node.Right})
			}
		}

		for parent, num := range curr {
			if parent == nil {
				root.Val = 0
			} else {
				if parent.Left != nil {
					parent.Left.Val = sum - num
				}

				if parent.Right != nil {
					parent.Right.Val = sum - num
				}
			}
		}

		queue = queue[qLen:]
	}

	return root
}

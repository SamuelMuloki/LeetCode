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

type GrandNode struct {
	node            *utils.TreeNode
	parent, gParent int
}

func SumEvenGrandparent(root *utils.TreeNode) int {
	queue := make([]*GrandNode, 0)
	queue = append(queue, &GrandNode{root, -1, -1})

	var ans int
	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			if queue[i].gParent%2 == 0 {
				ans += queue[i].node.Val
			}

			if queue[i].node.Left != nil {
				queue = append(queue, &GrandNode{
					node:    queue[i].node.Left,
					parent:  queue[i].node.Val,
					gParent: queue[i].parent,
				})
			}

			if queue[i].node.Right != nil {
				queue = append(queue, &GrandNode{
					node:    queue[i].node.Right,
					parent:  queue[i].node.Val,
					gParent: queue[i].parent,
				})
			}
		}

		queue = queue[qLen:]
	}

	return ans
}

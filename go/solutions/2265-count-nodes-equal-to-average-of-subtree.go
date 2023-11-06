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

type PairNode struct {
	sum, count int
}

func AverageOfSubtree(root *utils.TreeNode) int {
	count := 0
	postOrder(root, &count)

	return count
}

func postOrder(root *utils.TreeNode, count *int) PairNode {
	if root == nil {
		return PairNode{0, 0}
	}

	left := postOrder(root.Left, count)
	right := postOrder(root.Right, count)

	nodeSum := left.sum + right.sum + root.Val
	nodeCount := left.count + right.count + 1

	if root.Val == nodeSum/nodeCount {
		*count++
	}

	return PairNode{nodeSum, nodeCount}
}

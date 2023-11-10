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
func SumEvenGrandparent(root *utils.TreeNode) int {
	return findSum(root, -1, -1)
}

func findSum(root *utils.TreeNode, parent, gParent int) int {
	if root == nil {
		return 0
	}

	val := 0
	if gParent%2 == 0 {
		val = root.Val
	}

	return findSum(root.Left, root.Val, parent) + findSum(root.Right, root.Val, parent) + val
}

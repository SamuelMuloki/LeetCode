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
func ConstructFromPrePost(preorder []int, postorder []int) *utils.TreeNode {
	var constructTree func(preIndex, postIndex *int) *utils.TreeNode
	constructTree = func(preIndex, postIndex *int) *utils.TreeNode {
		node := &utils.TreeNode{Val: preorder[*preIndex]}
		*preIndex++
		if node.Val != postorder[*postIndex] {
			node.Left = constructTree(preIndex, postIndex)
		}

		if node.Val != postorder[*postIndex] {
			node.Right = constructTree(preIndex, postIndex)
		}

		*postIndex++

		return node
	}

	preIndex, postIndex := 0, 0
	return constructTree(&preIndex, &postIndex)
}

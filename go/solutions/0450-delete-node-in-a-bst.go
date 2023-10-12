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
func DeleteNode(root *utils.TreeNode, key int) *utils.TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}

		if root.Left == nil && root.Right != nil {
			return root.Right
		}

		if root.Left != nil && root.Right == nil {
			return root.Left
		}

		if root.Left != nil && root.Right != nil {
			minRight := findMin(root.Right)
			leftTree := root.Left
			rightTree := DeleteNode(root.Right, minRight.Val)
			minRight.Left = leftTree
			minRight.Right = rightTree
			return minRight
		}
	}

	if root.Val > key {
		root.Left = DeleteNode(root.Left, key)
	}

	if root.Val < key {
		root.Right = DeleteNode(root.Right, key)
	}

	return root
}

func findMin(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		return root
	}

	return findMin(root.Left)
}

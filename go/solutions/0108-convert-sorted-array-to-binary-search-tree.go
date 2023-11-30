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
func SortedArrayToBST(nums []int) *utils.TreeNode {
	var construct func(root *utils.TreeNode, arr []int)
	construct = func(root *utils.TreeNode, arr []int) {
		if root == nil {
			return
		}

		mid := len(arr) / 2
		root.Val = arr[mid]
		if len(arr[:mid]) > 0 {
			root.Left = &utils.TreeNode{}
			construct(root.Left, arr[:mid])
		}

		if len(arr[mid+1:]) > 0 {
			root.Right = &utils.TreeNode{}
			construct(root.Right, arr[mid+1:])
		}
	}

	root := &utils.TreeNode{}
	construct(root, nums)

	return root
}

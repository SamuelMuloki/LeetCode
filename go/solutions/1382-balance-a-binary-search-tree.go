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
func BalanceBST(root *utils.TreeNode) *utils.TreeNode {
	arr := make([]int, 0)
	var dfs func(node *utils.TreeNode)
	dfs = func(node *utils.TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		arr = append(arr, node.Val)
		dfs(node.Right)
	}

	dfs(root)

	var construct func(nums []int) *utils.TreeNode
	construct = func(nums []int) *utils.TreeNode {
		n := len(nums)
		if n == 0 {
			return nil
		}

		mid := n / 2
		return &utils.TreeNode{
			Val:   nums[mid],
			Left:  construct(nums[:mid]),
			Right: construct(nums[mid+1:]),
		}
	}

	return construct(arr)
}

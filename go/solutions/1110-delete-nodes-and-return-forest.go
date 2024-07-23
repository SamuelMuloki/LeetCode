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
func DelNodes(root *utils.TreeNode, to_delete []int) []*utils.TreeNode {
	deleteSet := make(map[int]bool)
	for _, val := range to_delete {
		deleteSet[val] = true
	}

	var ans []*utils.TreeNode
	var dfs func(node *utils.TreeNode) *utils.TreeNode
	dfs = func(node *utils.TreeNode) *utils.TreeNode {
		if node == nil {
			return nil
		}

		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)

		if _, ok := deleteSet[node.Val]; ok {
			if node.Left != nil {
				ans = append(ans, node.Left)
			}

			if node.Right != nil {
				ans = append(ans, node.Right)
			}

			return nil
		}

		return node
	}

	root = dfs(root)
	if root != nil {
		ans = append(ans, root)
	}

	return ans
}

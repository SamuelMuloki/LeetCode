// https://leetcode.com/problems/diameter-of-binary-tree/
package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func DiameterOfBinaryTree(root *utils.TreeNode) int {
	maxDiameter := 0
	dfs(root, &maxDiameter)

	return maxDiameter
}

func dfs(root *utils.TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}

	leftDiameter := dfs(root.Left, maxDiameter)
	rightDiameter := dfs(root.Right, maxDiameter)

	*maxDiameter = max(*maxDiameter, leftDiameter+rightDiameter)

	return max(leftDiameter, rightDiameter) + 1
}

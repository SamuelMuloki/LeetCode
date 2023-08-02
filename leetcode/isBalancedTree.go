// https://leetcode.com/problems/balanced-binary-tree/
package leetcode

func IsBalanced(root *TreeNode) bool {
	isValid := true
	findDepth(root, &isValid)
	return isValid
}

func findDepth(root *TreeNode, isValid *bool) int {
	if root == nil {
		return 0
	}

	leftTreeDepth := findDepth(root.Left, isValid)
	rightTreeDepth := findDepth(root.Right, isValid)

	minDepth, maxDepth := minMax(leftTreeDepth, rightTreeDepth)
	if *isValid && (maxDepth-minDepth) > 1 {
		*isValid = false
	}

	return maxDepth + 1
}

func minMax(x int, y int) (int, int) {
	if x > y {
		return y, x
	}

	return x, y
}

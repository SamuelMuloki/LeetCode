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
	n := len(nums)
	if n == 0 {
		return nil
	}

	mid := n / 2
	return &utils.TreeNode{
		Val:   nums[mid],
		Left:  SortedArrayToBST(nums[:mid]),
		Right: SortedArrayToBST(nums[mid+1:]),
	}
}

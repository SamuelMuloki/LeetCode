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
func ConstructMaximumBinaryTree(nums []int) *utils.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	index := 0
	for i := range nums {
		if nums[index] < nums[i] {
			index = i
		}
	}

	left := ConstructMaximumBinaryTree(nums[0:index])
	right := ConstructMaximumBinaryTree(nums[index+1:])

	return &utils.TreeNode{Val: nums[index], Left: left, Right: right}
}

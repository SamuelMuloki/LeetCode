package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func SortedListToBST(head *utils.ListNode) *utils.TreeNode {
	arr, curr := []int{}, head
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}

	var sortedArrToBST func(nums []int) *utils.TreeNode
	sortedArrToBST = func(nums []int) *utils.TreeNode {
		n := len(nums)
		if n == 0 {
			return nil
		}

		mid := n / 2
		return &utils.TreeNode{
			Val:   nums[mid],
			Left:  sortedArrToBST(nums[:mid]),
			Right: sortedArrToBST(nums[mid+1:]),
		}
	}

	return sortedArrToBST(arr)
}

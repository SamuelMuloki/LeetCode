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
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &utils.TreeNode{Val: head.Val}
	}

	prev, slow, fast := new(utils.ListNode), head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil
	return &utils.TreeNode{
		Val:   slow.Val,
		Left:  SortedListToBST(head),
		Right: SortedListToBST(slow.Next),
	}
}

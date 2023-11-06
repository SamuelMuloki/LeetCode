package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DeleteNode2(node *utils.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

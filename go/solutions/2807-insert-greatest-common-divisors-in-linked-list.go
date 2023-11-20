package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func InsertGreatestCommonDivisors(head *utils.ListNode) *utils.ListNode {
	curr := head
	var prev *utils.ListNode
	for curr != nil {
		if prev != nil {
			num := gcd(prev.Val, curr.Val)
			prev.Next = &utils.ListNode{Val: num, Next: curr}
		}

		prev = curr
		curr = curr.Next
	}

	return head
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}

	return gcd(y, x%y)
}

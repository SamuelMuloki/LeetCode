package solutions

import (
	"github.com/SamuelMuloki/LeetCode/go/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func GetDecimalValue(head *utils.ListNode) int {
	ans := 0
	for head != nil {
		ans *= 2
		ans += head.Val
		head = head.Next
	}

	return ans
}

package solutions

import (
	"strconv"

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
	curr, str := head, ""
	for curr != nil {
		str += strconv.Itoa(curr.Val)
		curr = curr.Next
	}

	ans, _ := strconv.ParseInt(str, 2, 32)
	return int(ans)
}

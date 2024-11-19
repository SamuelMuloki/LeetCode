package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SplitListToParts(head *utils.ListNode, k int) []*utils.ListNode {
	length := 0
	current := head
	var parts []*utils.ListNode

	for current != nil {
		length++
		current = current.Next
	}

	baseSize, extra := length/k, length%k
	current = head

	for i := 0; i < k; i++ {
		partSize := baseSize
		if extra > 0 {
			partSize++
			extra--
		}

		var partHead, partTail *utils.ListNode

		for j := 0; j < partSize; j++ {
			if partHead == nil {
				partHead, partTail = current, current
			} else {
				partTail.Next = current
				partTail = partTail.Next
			}

			if current != nil {
				current = current.Next
			}
		}

		if partTail != nil {
			partTail.Next = nil
		}

		parts = append(parts, partHead)
	}

	return parts
}

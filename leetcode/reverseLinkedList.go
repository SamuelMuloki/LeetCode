package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head ListNode
	Tail ListNode
}

func ReverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}

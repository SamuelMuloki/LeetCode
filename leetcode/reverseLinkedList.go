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
	if head == nil {
		return nil
	}

	curr := head
	var prev *ListNode
	for curr.Next != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	head = curr
	head.Next = prev

	return head
}

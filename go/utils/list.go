package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) PrintList() {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	node.Next.PrintList()
}

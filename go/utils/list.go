package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	var list []*ListNode
	for i := 0; i < len(arr); i++ {
		list = append(list, &ListNode{Val: arr[i]})
	}

	for i := 1; i < len(list); i++ {
		list[i-1].Next = list[i]
	}

	list[len(list)-1].Next = nil

	return list[0]
}

func (node *ListNode) PrintList() {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	node.Next.PrintList()
}

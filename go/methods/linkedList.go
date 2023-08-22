package methods

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

func (l LinkedList) Display() {
	for l.Head != nil {
		fmt.Printf("%d -> ", l.Head.Data)
		l.Head = l.Head.Next
	}
	fmt.Println()
}

func (l LinkedList) Front() (int, error) {
	if l.Head == nil {
		return 0, fmt.Errorf("Cannot find Front value in an empty Linked list")
	}
	return l.Head.Data, nil
}

func (l LinkedList) Back() (int, error) {
	if l.Head == nil {
		return 0, fmt.Errorf("Cannot find Back value in an empty Linked list")
	}
	return l.Tail.Data, nil
}

func (l *LinkedList) PushBack(node *Node) {
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		l.Length++
	} else {
		l.Tail.Next = node
		l.Tail = node
		l.Length++
	}
}

func (l *LinkedList) Delete(key int) {
	if l.Head.Data == key {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	var prev *Node = nil
	curr := l.Head
	for curr != nil && curr.Data != key {
		prev = curr
		curr = curr.Next
	}
	if curr == nil {
		fmt.Println("Key not found")
		return
	}
	prev.Next = curr.Next
	l.Length--
	fmt.Println("Node Deleted")
}

func (l *LinkedList) Reverse() {
	curr := l.Head
	l.Tail = l.Head
	var prev *Node
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	l.Head = prev
}

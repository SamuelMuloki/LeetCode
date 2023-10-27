package solutions

type Node struct {
	Val    int
	MinVal int
	Next   *Node
}

func NewNode(val, minVal int, next *Node) *Node {
	return &Node{val, minVal, next}
}

type MinStack struct {
	head *Node
}

func MinStackConstructor() MinStack {
	return MinStack{nil}
}

func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = NewNode(val, val, nil)
	} else {
		minVal := this.head.MinVal
		if val < minVal {
			minVal = val
		}
		this.head = NewNode(val, minVal, this.head)
	}
}

func (this *MinStack) Pop() {
	this.head = this.head.Next
}

func (this *MinStack) Top() int {
	return this.head.Val
}

func (this *MinStack) GetMin() int {
	return this.head.MinVal
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

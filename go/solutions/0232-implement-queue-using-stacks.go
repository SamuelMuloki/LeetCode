package solutions

type MyQueue struct {
	input  []int
	output []int
}

func MyQueueConstructor() MyQueue {
	return MyQueue{
		input:  []int{},
		output: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

func (this *MyQueue) Pop() int {
	ret := this.Peek()
	this.output = this.output[:len(this.output)-1]
	return ret
}

func (this *MyQueue) Peek() int {
	if len(this.output) != 0 {
		return this.output[len(this.output)-1]
	}

	for len(this.input) != 0 {
		this.output = append(this.output, this.input[len(this.input)-1])
		this.input = this.input[:len(this.input)-1]
	}

	return this.output[len(this.output)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.input) == 0 && len(this.output) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

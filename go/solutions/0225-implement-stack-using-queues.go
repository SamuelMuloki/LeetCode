package solutions

type MyStack struct {
	output []int
}

func MyStackConstructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.output = append(this.output, x)
	for i := 1; i < len(this.output); i++ {
		this.output = append(this.output, this.output[0])
		this.output = this.output[1:]
	}
}

func (this *MyStack) Pop() int {
	ret := this.Top()
	this.output = this.output[1:]
	return ret
}

func (this *MyStack) Top() int {
	return this.output[0]
}

func (this *MyStack) Empty() bool {
	return len(this.output) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

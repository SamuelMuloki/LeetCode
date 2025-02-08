package solutions

import "container/heap"

type NumberContainers struct {
	nums    map[int]*MinHeap
	indices map[int]int
}

func NumberContainersConstructor() NumberContainers {
	return NumberContainers{
		nums:    make(map[int]*MinHeap),
		indices: make(map[int]int),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if _, ok := this.nums[number]; !ok {
		h := &MinHeap{}
		heap.Init(h)
		this.nums[number] = h
	}

	heap.Push(this.nums[number], index)
	this.indices[index] = number
}

func (this *NumberContainers) Find(number int) int {
	if _, ok := this.nums[number]; !ok {
		return -1
	}

	for this.nums[number].Len() > 0 {
		top := (*this.nums[number])[0]
		if this.indices[top] == number {
			break
		}

		heap.Pop(this.nums[number])
	}

	if this.nums[number].Len() > 0 {
		return (*this.nums[number])[0]
	}

	return -1
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

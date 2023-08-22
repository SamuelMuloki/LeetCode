package methods

import "fmt"

type Queue struct {
	Elements []int
	Size     int
}

func (q *Queue) Enqueue(element int) {
	if q.Length() == q.Size {
		fmt.Println("Overflow")
		return
	}
	q.Elements = append(q.Elements, element)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("Underflow")
		return 0
	}

	element := q.Elements[0]
	if q.Length() == 1 {
		q.Elements = nil
		return element
	}

	q.Elements = q.Elements[1:]
	return element
}

func (q Queue) Length() int {
	return len(q.Elements)
}

func (q Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Empty Queue")
	}

	return q.Elements[0], nil
}

package solutions

import "container/heap"

func LongestDiverseString(a int, b int, c int) string {
	pq := &CharCountMaxHeap{}
	heap.Init(pq)
	if a > 0 {
		heap.Push(pq, CharCount{a, 'a'})
	}

	if b > 0 {
		heap.Push(pq, CharCount{b, 'b'})
	}

	if c > 0 {
		heap.Push(pq, CharCount{c, 'c'})
	}

	ans := ""
	for pq.Len() > 0 {
		first := heap.Pop(pq).(CharCount)
		if len(ans) >= 2 && ans[len(ans)-1] == first.ch && ans[len(ans)-2] == first.ch {
			if pq.Len() == 0 {
				break
			}

			second := heap.Pop(pq).(CharCount)
			ans += string(second.ch)
			second.count--
			if second.count > 0 {
				heap.Push(pq, second)
			}
			heap.Push(pq, first)
		} else {
			ans += string(first.ch)
			first.count--
			if first.count > 0 {
				heap.Push(pq, first)
			}
		}
	}

	return ans
}

type CharCount struct {
	count int
	ch    byte
}

type CharCountMaxHeap []CharCount

func (h CharCountMaxHeap) Len() int           { return len(h) }
func (h CharCountMaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h CharCountMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *CharCountMaxHeap) Push(x any)        { *h = append(*h, x.(CharCount)) }
func (h *CharCountMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

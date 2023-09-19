package methods

type node struct {
	v    *Vertex
	next *node
}

type queue struct {
	head *node
	tail *node
}

func (q *queue) enqueue(v *Vertex) {
	n := &node{v: v}

	if q.tail == nil {
		q.head = n
		q.tail = n
		return
	}

	q.tail.next = n
	q.tail = n
}

func (q *queue) dequeue() *Vertex {
	n := q.head
	if n == nil {
		return nil
	}

	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}

	return n.v
}

func (g *Graph) BFS(start *Vertex, visitCb func(int)) {
	vertexQueue := &queue{}
	visited := make(map[int]bool)

	curr := start
	for {
		visitCb(curr.Key)
		visited[curr.Key] = true

		for _, v := range curr.Vertices {
			if !visited[v.Key] {
				vertexQueue.enqueue(v)
			}
		}

		curr = vertexQueue.dequeue()
		if curr == nil {
			break
		}
	}
}

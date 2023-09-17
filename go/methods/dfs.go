package methods

func (g *Graph) DFS(start *Vertex, visitCb func(int)) {
	visited := make(map[int]bool)

	if start == nil {
		return
	}

	visited[start.Key] = true
	visitCb(start.Key)

	for _, v := range start.Vertices {
		if !visited[v.Key] {
			g.DFS(v, visitCb)
		}
	}
}

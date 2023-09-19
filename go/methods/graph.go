package methods

import "fmt"

type Vertex struct {
	Key      int
	Vertices map[int]*Vertex
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: make(map[int]*Vertex),
	}
}

func (v *Vertex) ToString() string {
	s := fmt.Sprintf("%d:", v.Key)

	for _, neigbor := range v.Vertices {
		s = fmt.Sprintf(" %d", neigbor.Key)
	}

	return s
}

type Graph struct {
	Vertices map[int]*Vertex
	directed bool
}

func NewDirectedGraph() *Graph {
	return &Graph{
		Vertices: make(map[int]*Vertex),
		directed: true,
	}
}

func NewUnDirectedGraph() *Graph {
	return &Graph{
		Vertices: make(map[int]*Vertex),
	}
}

func (g *Graph) AddVertex(key int) {
	v := NewVertex(key)
	g.Vertices[key] = v
}

func (g *Graph) AddEdge(v1, v2 int) {
	eV1, eV2 := g.Vertices[v1], g.Vertices[v2]

	if eV1 == nil || eV2 == nil {
		panic("one or both of the vertices don't exist")
	}

	if eV1.Vertices[eV2.Key] != nil {
		return
	}

	eV1.Vertices[eV2.Key] = eV2
	if !g.directed && eV1.Key != eV2.Key {
		eV2.Vertices[eV1.Key] = eV1
	}

	g.Vertices[eV1.Key] = eV1
	g.Vertices[eV2.Key] = eV2
}

func (g *Graph) String() string {
	s := ""
	i := 0
	for _, v := range g.Vertices {
		if i != 0 {
			s += "\n"
		}
		s += v.ToString()
		i++
	}
	return s
}

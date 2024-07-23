package solutions

type MaxUnionFind struct {
	Parents []int
	Ranks   []int
	Groups  int
}

func NewMaxUnionFind(n int) *MaxUnionFind {
	parents := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parents[i] = i
	}

	return &MaxUnionFind{
		Parents: parents,
		Ranks:   make([]int, n+1),
		Groups:  n,
	}
}

func MaxNumEdgesToRemove(n int, edges [][]int) int {
	AliceUF := NewMaxUnionFind(n)
	BobUF := NewMaxUnionFind(n)

	edgesRequired := 0
	for _, edge := range edges {
		if edge[0] == 3 {
			res1 := AliceUF.PerformUnion(edge)
			res2 := BobUF.PerformUnion(edge)
			if res1 || res2 {
				edgesRequired++
			}
		}
	}

	for _, edge := range edges {
		if edge[0] == 1 && AliceUF.PerformUnion(edge) {
			edgesRequired++
		} else if edge[0] == 2 && BobUF.PerformUnion(edge) {
			edgesRequired++
		}
	}

	if AliceUF.isConnected() && BobUF.isConnected() {
		return len(edges) - edgesRequired
	}

	return -1
}

func (uf *MaxUnionFind) PerformUnion(edge []int) bool {
	p1 := uf.findParent(edge[1])
	p2 := uf.findParent(edge[2])

	if p1 == p2 {
		return false
	}

	if uf.Ranks[p1] > uf.Ranks[p2] {
		uf.Parents[p2] = p1
	} else if uf.Ranks[p2] > uf.Ranks[p1] {
		uf.Parents[p1] = p2
	} else {
		uf.Parents[p2] = p1
		uf.Ranks[p1]++
	}

	uf.Groups--
	return true
}

func (uf *MaxUnionFind) findParent(idx int) int {
	if idx != uf.Parents[idx] {
		uf.Parents[idx] = uf.findParent(uf.Parents[idx])
	}
	return uf.Parents[idx]
}

func (uf *MaxUnionFind) isConnected() bool {
	return uf.Groups == 1
}

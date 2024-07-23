package solutions

type DisjointSet struct {
	roots []int
	ranks []int
}

func newDisjointSet(n int) *DisjointSet {
	roots := make([]int, n)
	ranks := make([]int, n)
	for i := 0; i < n; i++ {
		roots[i] = i
		ranks[i] = 1
	}

	newDisjointSet := DisjointSet{roots, ranks}
	return &newDisjointSet
}

func (ds *DisjointSet) Find(x int) int {
	if ds.roots[x] == x {
		return x
	}

	ds.roots[x] = ds.Find(ds.roots[x])
	return ds.roots[x]
}

func (ds *DisjointSet) Union(x, y int) {
	rootX := ds.Find(x)
	rootY := ds.Find(y)
	if rootX == rootY {
		return
	}

	rankX := ds.ranks[rootX]
	rankY := ds.ranks[rootY]
	if rankX > rankY {
		ds.roots[rootY] = rootX
		return
	}

	if rankX < rankY {
		ds.roots[rootX] = rootY
		return
	}

	ds.roots[rootX] = rootY
	ds.ranks[rootY] += 1
}

func (ds *DisjointSet) IsConnected(x, y int) bool {
	return ds.Find(x) == ds.Find(y)
}

func ValidPath(n int, edges [][]int, source int, destination int) bool {
	ds := newDisjointSet(n)
	for _, edge := range edges {
		ds.Union(edge[0], edge[1])
	}

	return ds.IsConnected(source, destination)
}

package solutions

type DSU struct {
	parent, rank, size []int
}

func NewDSU(n int) *DSU {
	parent, size := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &DSU{parent, make([]int, n), size}
}

func (ds *DSU) Find(node int) int {
	if ds.parent[node] == node {
		return ds.parent[node]
	}

	return ds.Find(ds.parent[node])
}

func (ds *DSU) Union(x, y int) {
	rootX, rootY := ds.Find(x), ds.Find(y)
	if rootX == rootY {
		return
	}

	if rootX < rootY {
		rootX, rootY = rootY, rootX
	}

	ds.parent[rootY] = rootX
	ds.size[rootX] += ds.size[rootY]
	if ds.rank[rootX] == ds.rank[rootY] {
		ds.rank[rootX]++
	}
}

func CountCompleteComponents(n int, edges [][]int) int {
	ds := NewDSU(n)
	degree := make([]int, n)
	for _, edge := range edges {
		ds.Union(edge[0], edge[1])
		degree[edge[0]]++
		degree[edge[1]]++
	}

	ans := 0
	seen := make(map[int]bool)
	for i := 0; i < n; i++ {
		root := ds.Find(i)
		if seen[root] {
			continue
		}

		seen[root] = true
		size := ds.size[root]
		isComplete := true
		for j := 0; j < n; j++ {
			if ds.Find(j) == root && degree[j] != size-1 {
				isComplete = false
				break
			}
		}
		if isComplete {
			ans++
		}
	}

	return ans
}

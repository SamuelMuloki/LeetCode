package solutions

type DSU2 struct {
	parent []int
}

func NewDSU2(n int) *DSU2 {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &DSU2{parent}
}

func (ds *DSU2) Find(node int) int {
	if ds.parent[node] == node {
		return ds.parent[node]
	}

	return ds.Find(ds.parent[node])
}

func (ds *DSU2) Union(x, y int) {
	rootX, rootY := ds.Find(x), ds.Find(y)
	if rootX == rootY {
		return
	}

	if rootX < rootY {
		ds.parent[rootY] = rootX
	} else {
		ds.parent[rootX] = rootY
	}

}

func SmallestEquivalentString(s1 string, s2 string, baseStr string) string {
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}

	ds := NewDSU2(26)
	for i := 0; i < len(s1); i++ {
		if i > len(s2) {
			break
		}

		ds.Union(int(s1[i]-'a'), int(s2[i]-'a'))
	}

	res := []byte(baseStr)
	for i := 0; i < len(res); i++ {
		val := ds.Find(int(res[i] - 'a'))
		res[i] = min(res[i], byte(val+'a'))
	}

	return string(res)
}

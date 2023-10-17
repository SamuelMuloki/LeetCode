package solutions

type UnionFind struct {
	n, components int
	parents       []int
}

func newUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}

	return &UnionFind{n, n, parents}
}

func (uf *UnionFind) Union(parent, child int) bool {
	parentParent := uf.Find(parent)
	childParent := uf.Find(child)

	if childParent != child || parentParent == childParent {
		return false
	}

	uf.components--
	uf.parents[childParent] = parentParent
	return true
}

func (uf *UnionFind) Find(node int) int {
	if uf.parents[node] != node {
		uf.parents[node] = uf.Find(uf.parents[node])
	}

	return uf.parents[node]
}

func ValidateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	uf := newUnionFind(n)
	for node := 0; node < n; node++ {
		children := []int{leftChild[node], rightChild[node]}
		for _, child := range children {
			if child == -1 {
				continue
			}

			if !uf.Union(node, child) {
				return false
			}
		}
	}

	return uf.components == 1
}

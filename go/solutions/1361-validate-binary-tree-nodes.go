package solutions

func ValidateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	findRoot := func() int {
		children := make(map[int]bool)

		for i := range leftChild {
			children[leftChild[i]] = true
		}

		for i := range rightChild {
			children[rightChild[i]] = true
		}

		for i := 0; i < n; i++ {
			if !children[i] {
				return i
			}
		}

		return -1
	}

	root := findRoot()
	if root == -1 {
		return false
	}

	seen := make(map[int]bool)
	st := make([]int, 0)
	seen[root] = true
	st = append(st, root)

	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]

		children := []int{leftChild[node], rightChild[node]}
		for _, child := range children {
			if child != -1 {
				if seen[child] {
					return false
				}

				seen[child] = true
				st = append(st, child)
			}
		}
	}

	return len(seen) == n
}

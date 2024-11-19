package solutions

func RemoveStones(stones [][]int) int {
	componentCount := 0
	uniqueNodes := make(map[int]int)

	var find func(int) int
	find = func(element int) int {
		if _, exists := uniqueNodes[element]; !exists {
			uniqueNodes[element] = element
			componentCount++
		}
		if uniqueNodes[element] != element {
			uniqueNodes[element] = find(uniqueNodes[element])
		}
		return uniqueNodes[element]
	}

	union := func(elementA, elementB int) {
		root1 := find(elementA)
		root2 := find(elementB)
		if root1 != root2 {
			uniqueNodes[root2] = root1
			componentCount--
		}
	}

	for _, stone := range stones {
		union(stone[0]+1, stone[1]+10002)
	}

	return len(stones) - componentCount
}

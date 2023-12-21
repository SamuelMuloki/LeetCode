package solutions

func CanVisitAllRooms(rooms [][]int) bool {
	seen := make([]bool, len(rooms))
	seen[0] = true
	st := make([]int, 0)
	st = append(st, 0)

	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		for _, nei := range rooms[node] {
			if !seen[nei] {
				seen[nei] = true
				st = append(st, nei)
			}
		}
	}

	for _, val := range seen {
		if !val {
			return false
		}
	}

	return true
}

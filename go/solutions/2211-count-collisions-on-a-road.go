package solutions

func CountCollisions(directions string) int {
	i := 0
	n := len(directions)
	for i < n && directions[i] == 'L' {
		i++
	}

	j := n - 1
	for j >= 0 && directions[j] == 'R' {
		j--
	}

	res := 0
	for k := i; k <= j; k++ {
		if directions[k] != 'S' {
			res++
		}
	}

	return res
}

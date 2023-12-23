package solutions

type Point struct {
	x, y int
}

func IsPathCrossing(path string) bool {
	point := Point{0, 0}
	paths := map[Point]bool{point: true}
	for i := range path {
		if path[i] == 'N' {
			point.y++
		} else if path[i] == 'E' {
			point.x++
		} else if path[i] == 'W' {
			point.x--
		} else if path[i] == 'S' {
			point.y--
		}

		if paths[point] {
			return paths[point]
		}

		paths[point] = true
	}

	return false
}

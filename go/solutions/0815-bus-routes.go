package solutions

func NumBusesToDestination(routes [][]int, source int, target int) int {
	var n = len(routes)

	var m = make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < len(routes[i]); j++ {
			var busstop = routes[i][j]
			var buses = m[busstop]
			buses = append(buses, i)
			m[busstop] = buses
		}
	}

	var mbuses = make(map[int]bool)
	var mstops = make(map[int]bool)

	var queue []int
	var count = 0

	queue = append(queue, source)
	mstops[source] = true

	for len(queue) > 0 {
		var qLen = len(queue)

		for i := 0; i < qLen; i++ {
			var busstop = queue[0]
			queue = queue[1:]
			if busstop == target {
				return count
			}

			var buses = m[busstop]

			for _, bus := range buses {
				if _, ok := mbuses[bus]; ok {
					continue
				}

				for _, stop := range routes[bus] {
					if _, ok := mstops[stop]; ok {
						continue
					}

					queue = append(queue, stop)
					mstops[stop] = true
				}

				mbuses[bus] = true
			}

		}

		count += 1
	}

	return -1
}

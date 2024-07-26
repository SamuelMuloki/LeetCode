package solutions

func FindTheCity(n int, edges [][]int, distanceThreshold int) int {
	INF := 1000000007
	distanceMatrix := make([][]int, n)
	for i := range distanceMatrix {
		distanceMatrix[i] = make([]int, n)
		for j := range distanceMatrix[i] {
			distanceMatrix[i][j] = INF
		}
		distanceMatrix[i][i] = 0
	}

	for _, edge := range edges {
		start, end, weight := edge[0], edge[1], edge[2]
		distanceMatrix[start][end] = weight
		distanceMatrix[end][start] = weight
	}

	var floyd = func() {
		for k := 0; k < n; k++ {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					distanceMatrix[i][j] = min(
						distanceMatrix[i][j],
						distanceMatrix[i][k]+distanceMatrix[k][j],
					)
				}
			}
		}
	}

	var getCityWithFewestReachable = func() int {
		cityWithFewestReachable, fewestReachableCount := -1, n
		for i := 0; i < n; i++ {
			reachableCount := 0
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				if distanceMatrix[i][j] <= distanceThreshold {
					reachableCount++
				}
			}
			if reachableCount <= fewestReachableCount {
				fewestReachableCount = reachableCount
				cityWithFewestReachable = i
			}
		}

		return cityWithFewestReachable
	}

	floyd()

	return getCityWithFewestReachable()
}

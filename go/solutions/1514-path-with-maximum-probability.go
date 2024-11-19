package solutions

func MaxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	maxProb := []float64{}
	for i := 0; i < n; i++ {
		maxProb = append(maxProb, float64(0))
	}
	maxProb[start_node] = float64(1)

	for i := 0; i < n-1; i++ {
		hasUpdate := false
		for j := 0; j < len(edges); j++ {
			u, v := edges[j][0], edges[j][1]
			pathProb := succProb[j]
			if maxProb[u]*pathProb > maxProb[v] {
				maxProb[v] = maxProb[u] * pathProb
				hasUpdate = true
			}

			if maxProb[v]*pathProb > maxProb[u] {
				maxProb[u] = maxProb[v] * pathProb
				hasUpdate = true
			}
		}

		if !hasUpdate {
			break
		}
	}

	return maxProb[end_node]
}

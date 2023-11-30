package solutions

func FindCircleNum(isConnected [][]int) int {
	n, numberOfComponents := len(isConnected), 0
	visit := make([]bool, n)

	var dfs func(node int)
	dfs = func(node int) {
		visit[node] = true
		for i := 0; i < n; i++ {
			if isConnected[node][i] == 1 && !visit[i] {
				dfs(i)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visit[i] {
			numberOfComponents++
			dfs(i)
		}
	}

	return numberOfComponents
}

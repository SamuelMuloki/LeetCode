package solutions

import "math"

func KthGrammar(n int, k int) int {
	var dfs func(n, k, rootVal int) int
	dfs = func(n, k, rootVal int) int {
		if n == 1 {
			return rootVal
		}

		totalNodes := math.Pow(float64(2), float64(n-1))
		halfNodes := int(totalNodes) / 2

		if k > halfNodes {
			nextRootVal := 0
			if rootVal == 0 {
				nextRootVal = 1
			}

			return dfs(n-1, k-halfNodes, nextRootVal)
		} else {
			nextRootVal := 1
			if rootVal == 0 {
				nextRootVal = 0
			}

			return dfs(n-1, k, nextRootVal)
		}
	}

	return dfs(n, k, 0)
}

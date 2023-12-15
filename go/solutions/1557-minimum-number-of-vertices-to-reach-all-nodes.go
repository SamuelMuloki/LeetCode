package solutions

func FindSmallestSetOfVertices(n int, edges [][]int) []int {
	ans, seen := make([]int, 0), make([]int, n)
	for i := range edges {
		seen[edges[i][1]] = 1
	}

	for i := 0; i < n; i++ {
		if seen[i] == 0 {
			ans = append(ans, i)
		}
	}

	return ans
}

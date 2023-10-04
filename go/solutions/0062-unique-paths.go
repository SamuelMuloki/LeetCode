package solutions

// (m-1 + n-1)!/(m-1)!(n-1)! -> (m-1) Down moves, (n-1) right moves
func UniquePaths(m int, n int) int {
	path := 1
	for i := n; i < (m + n - 1); i++ {
		path *= i
		path /= (i - n + 1)
	}

	return path
}

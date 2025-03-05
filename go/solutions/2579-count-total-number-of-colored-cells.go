package solutions

func ColoredCells(n int) int64 {
	return int64(1 + n*(n-1)*2)
}

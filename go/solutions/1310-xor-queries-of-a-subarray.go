package solutions

func XorQueries(arr []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i := range queries {
		xOr := 0
		for j := queries[i][0]; j <= queries[i][1]; j++ {
			xOr ^= arr[j]
		}
		ans[i] = xOr
	}
	return ans
}

package solutions

func XorQueries(arr []int, queries [][]int) []int {
	prefix := make([]int, len(arr)+1)
	prefix[0] = 0
	for i := 1; i <= len(arr); i++ {
		prefix[i] = prefix[i-1] ^ arr[i-1]
	}

	ans := make([]int, len(queries))
	for i := range queries {
		ans[i] = prefix[queries[i][1]+1] ^ prefix[queries[i][0]]
	}

	return ans
}

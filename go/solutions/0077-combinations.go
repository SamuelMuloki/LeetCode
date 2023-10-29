package solutions

func Combine(n int, k int) [][]int {
	ans := make([][]int, 0)

	var dfs func(num int, curr []int)
	dfs = func(num int, curr []int) {
		if len(curr) == k {
			ans = append(ans, append([]int{}, curr...))
			return
		}

		for i := num; i <= n; i++ {
			curr = append(curr, i)
			dfs(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(1, make([]int, 0))
	return ans
}

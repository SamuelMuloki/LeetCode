package solutions

func CountMaxOrSubsets(nums []int) int {
	n, targetOr := len(nums), 0
	for i := range nums {
		targetOr |= nums[i]
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, targetOr+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(idx, currOr int) int
	dfs = func(idx, currOr int) int {
		if idx == len(nums) {
			if currOr == targetOr {
				return 1
			}
			return 0
		}

		if memo[idx][currOr] != -1 {
			return memo[idx][currOr]
		}

		countWithout := dfs(idx+1, currOr)
		countWith := dfs(idx+1, currOr|nums[idx])

		memo[idx][currOr] = countWithout + countWith
		return memo[idx][currOr]
	}

	return dfs(0, 0)
}

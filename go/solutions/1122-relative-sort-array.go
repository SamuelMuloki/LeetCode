package solutions

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	memo := make([]int, 1001)
	for _, val1 := range arr1 {
		memo[val1]++
	}

	ans := []int{}
	for _, val2 := range arr2 {
		for memo[val2] > 0 {
			ans = append(ans, val2)
			memo[val2]--
		}
	}

	for i := range memo {
		for memo[i] > 0 {
			ans = append(ans, i)
			memo[i]--
		}
	}

	return ans
}

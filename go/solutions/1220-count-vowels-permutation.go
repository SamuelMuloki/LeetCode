package solutions

func CountVowelPermutation(n int) int {
	mod := 1000000007
	vowels := map[byte][]byte{
		's': {'a', 'e', 'i', 'o', 'u'},
		'a': {'e'},
		'e': {'a', 'i'},
		'i': {'a', 'e', 'o', 'u'},
		'o': {'i', 'u'},
		'u': {'a'},
	}

	dp := make(map[byte][]int)
	for i := range vowels {
		dp[i] = make([]int, n+1)
	}

	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(i int, vowel byte) int
	dfs = func(i int, vowel byte) int {
		if i == 0 {
			return 1
		}

		if dp[vowel][i] != -1 {
			return dp[vowel][i]
		}

		count := 0
		for _, v := range vowels[vowel] {
			count = (count + dfs(i-1, v)) % mod
		}

		dp[vowel][i] = count
		return dp[vowel][i]
	}

	return dfs(n, 's')
}

package solutions

func LongestPalindrome2(s string) int {
	count := make(map[byte]int)
	for i := range s {
		count[s[i]]++
	}

	ans, hasOdd := 0, false
	for _, num := range count {
		if num%2 == 0 {
			ans += num
		} else {
			hasOdd = true
			ans += num - 1
		}
	}

	if hasOdd {
		ans += 1
	}

	return ans
}

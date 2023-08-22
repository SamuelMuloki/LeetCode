// https://solutions.com/problems/palindrome-partitioning/
package solutions

func Partition(s string) [][]string {
	output := make([][]string, 0)
	curr := make([]string, 0)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == len(s) {
			output = append(output, append([]string{}, curr...))
			return
		}

		for i := idx; i < len(s); i++ {
			if isPalindrome(s[idx : i+1]) {
				curr = append(curr, s[idx:i+1])
				backtrack(i + 1)
				curr = curr[:len(curr)-1]
			}
		}
	}

	backtrack(0)

	return output
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

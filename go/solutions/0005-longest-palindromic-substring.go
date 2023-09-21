// https://leetcode.com/problems/longest-palindromic-substring/
package solutions

func LongestPalindrome(s string) string {
	max, maxStr := 1, string(s[0])
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			str := string(s[i]) + string(s[i+1:j+1])

			if isPalindrome(str) && len(str) > max {
				maxStr = str
				max = len(str)
			}
		}
	}

	return maxStr
}

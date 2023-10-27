// https://leetcode.com/problems/longest-palindromic-substring/
package solutions

func LongestPalindrome(s string) string {
	ans := []int{0, 0}

	for i := 0; i < len(s); i++ {
		oddLen := expand(i, i, s)
		if oddLen > ans[1]-ans[0]+1 {
			dist := oddLen / 2
			ans[0] = i - dist
			ans[1] = i + dist
		}

		evenLen := expand(i, i+1, s)
		if evenLen > ans[1]-ans[0]+1 {
			dist := (evenLen / 2) - 1
			ans[0] = i - dist
			ans[1] = i + 1 + dist
		}
	}

	i, j := ans[0], ans[1]
	return string(s[i : j+1])
}

func expand(i, j int, s string) int {
	l, r := i, j

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return r - l - 1
}

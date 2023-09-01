// https://leetcode.com/problems/palindrome-number
package solutions

func IsPalindromeNumber(x int) bool {
	reversed, temp := 0, x
	for temp > 0 {
		reversed = (reversed * 10) + (temp % 10)
		temp /= 10
	}

	return reversed == x
}

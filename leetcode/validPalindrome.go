package leetcode

import (
	"regexp"
	"strings"
)

/*
A phrase is a palindrome if, after converting all uppercase letters into
lowercase letters and removing all non-alphanumeric characters,
it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

E.g
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
*/

func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	nonAlphaNum := regexp.MustCompile("[^a-zA-Z0-9]+")
	str := strings.ToLower(nonAlphaNum.ReplaceAllString(s, ""))

	i, j := 0, len(str)-1
	for i <= j {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}

	return true
}

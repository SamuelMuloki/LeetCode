package leetcode

import "strings"

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different
word or phrase, typically using all the original letters exactly once.

Eg:
Input: s = "anagram", t = "nagaram"
Output: true

Input: s = "rat", t = "car"
Output: false
*/

func IsAnagram(s string, t string) bool {
	sCount := make(map[string]int)
	sArr := strings.Split(s, "")

	if len(s) != len(t) {
		return false
	}

	for _, val := range sArr {
		sCount[val] = sCount[val] + 1
	}

	for key, value := range sCount {
		if strings.Count(t, key) != value {
			return false
		}
	}
	return true
}

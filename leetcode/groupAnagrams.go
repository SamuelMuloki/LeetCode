package leetcode

import "sort"

/*
Given an array of strings strs, group the anagrams together.
You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a
different word or phrase, typically using all the original letters exactly once.

Eg:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Input: strs = [""]
Output: [[""]]

Input: strs = ["a"]
Output: [["a"]]
*/

func GroupAnagrams(strs []string) [][]string {
	anagrams := [][]string{}
	evaluated := make(map[string][]string)

	for i := range strs {
		key := sortString(strs[i])
		evaluated[key] = append(evaluated[key], strs[i])
	}

	for idx := range evaluated {
		anagrams = append(anagrams, evaluated[idx])
	}

	return anagrams
}

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(a, b int) bool {
		return runes[a] < runes[b]
	})
	return string(runes)
}

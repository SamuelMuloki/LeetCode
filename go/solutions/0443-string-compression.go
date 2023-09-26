// https://leetcode.com/problems/string-compression
package solutions

import "strconv"

func Compress(chars []byte) int {
	i, res := 0, 0
	for i < len(chars) {
		groupLength := 1
		for i+groupLength < len(chars) && chars[i+groupLength] == chars[i] {
			groupLength++
		}

		chars[res] = chars[i]
		res++
		if groupLength > 1 {
			str := strconv.Itoa(groupLength)
			for _, char := range str {
				chars[res] = byte(char)
				res++
			}
		}

		i += groupLength
	}

	return res
}

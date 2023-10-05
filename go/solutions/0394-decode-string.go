package solutions

import (
	"strconv"
)

func DecodeString(s string) string {
	countSt := make([]int, 0)
	strSt := make([]string, 0)
	var count int
	var str string

	for i := 0; i < len(s); i++ {
		curr := string(s[i])

		if n, err := strconv.Atoi(curr); err == nil {
			count = 10*count + n
		} else if curr == "[" {
			countSt = append(countSt, count)
			strSt = append(strSt, str)

			count = 0
			str = ""
		} else if curr == "]" {
			tempStr := strSt[len(strSt)-1]
			strSt = strSt[:len(strSt)-1]

			currCount := countSt[len(countSt)-1]
			countSt = countSt[:len(countSt)-1]

			for currCount > 0 {
				tempStr += str
				currCount--
			}
			str = tempStr
		} else {
			str += curr
		}
	}

	return str
}

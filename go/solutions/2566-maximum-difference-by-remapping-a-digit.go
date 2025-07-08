package solutions

import "strconv"

func MinMaxDifference(num int) int {
	str1 := []byte(strconv.Itoa(num))
	digit := str1[0]
	for i := 0; i < len(str1); i++ {
		if str1[i] != '9' {
			digit = str1[i]
			break
		}
	}

	for i := 0; i < len(str1); i++ {
		if str1[i] == digit {
			str1[i] = '9'
		}
	}

	str2 := []byte(strconv.Itoa(num))
	digit2 := str2[0]
	for i := 0; i < len(str2); i++ {
		if str2[i] == digit2 {
			str2[i] = '0'
		}
	}

	n1, _ := strconv.Atoi(string(str1))
	n2, _ := strconv.Atoi(string(str2))

	return n1 - n2
}

package solutions

import "strconv"

func MaxDiff(num int) int {
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
	j := 0
	for ; j < len(str2); j++ {
		if str2[j] == '1' || str2[j] == '0' {
			continue
		}
		break
	}

	if j < len(str2) {
		digit2 := str2[j]
		for i := 0; i < len(str2); i++ {
			if j > 0 && str2[i] == digit2 {
				str2[i] = '0'
			} else if str2[i] == digit2 {
				str2[i] = '1'
			}
		}
	}

	n1, _ := strconv.Atoi(string(str1))
	n2, _ := strconv.Atoi(string(str2))

	return n1 - n2
}

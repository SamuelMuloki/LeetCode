package solutions

import "strconv"

func Maximum69Number(num int) int {
	byteArr := []byte(strconv.Itoa(num))
	for i := 0; i < len(byteArr); i++ {
		if byteArr[i] != '9' {
			byteArr[i] = '9'
			break
		}
	}

	res, _ := strconv.Atoi(string(byteArr))
	return res
}

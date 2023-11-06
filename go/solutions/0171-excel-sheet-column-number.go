package solutions

import "math"

func TitleToNumber(columnTitle string) int {
	var num int
	j := len(columnTitle) - 1
	for i := 0; i < len(columnTitle); i++ {
		num += (int(columnTitle[i]-'A') + 1) * (int(math.Pow(26, float64(j))))
		j--
	}

	return num
}

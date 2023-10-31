package solutions

import (
	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func MaxNumberOfBalloons(text string) int {
	b, a, l, o, n := 0, 0, 0, 0, 0
	for _, r := range text {
		switch r {
		case 'b':
			b++
		case 'a':
			a++
		case 'l':
			l++
		case 'o':
			o++
		case 'n':
			n++
		}
	}

	ans := utils.Min(b, utils.Min(a, n))
	ans = utils.Min(ans, utils.Min(l/2, o/2))
	return ans
}

package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func IsReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	width, height := abs(sx-fx), abs(sy-fy)
	if width == 0 && height == 0 && t == 1 {
		return false
	}

	return t >= utils.Max(width, height)
}

package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func MaxArea(height []int) int {
	area := 0

	i, j := 0, len(height)-1
	for i < j {
		area = utils.Max(area, utils.Min(height[j], height[i])*(j-i))

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return int(area)
}

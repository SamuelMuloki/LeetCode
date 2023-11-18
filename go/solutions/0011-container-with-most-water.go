package solutions

func MaxArea(height []int) int {
	area := 0

	i, j := 0, len(height)-1
	for i < j {
		area = max(area, min(height[j], height[i])*(j-i))

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return int(area)
}

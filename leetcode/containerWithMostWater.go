package leetcode

func MaxArea(height []int) int {
	area := 0

	i, j := 0, len(height)-1
	for i < j {
		area = findMax(area, findMin(height[j], height[i])*(j-i))

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return int(area)
}

func findMin(x int, y int) int {
	if x < y {
		return x
	}

	return y
}

func findMax(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

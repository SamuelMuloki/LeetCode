package solutions

func QueryResults(limit int, queries [][]int) []int {
	ans := make([]int, len(queries))
	ballMap := make(map[int]int)
	colorMap := make(map[int]int)
	for i := 0; i < len(queries); i++ {
		ball, color := queries[i][0], queries[i][1]
		if prevColor, ok := ballMap[ball]; ok {
			colorMap[prevColor]--
			if colorMap[prevColor] == 0 {
				delete(colorMap, prevColor)
			}
		}

		ballMap[ball] = color
		colorMap[color]++
		ans[i] = len(colorMap)
	}

	return ans
}

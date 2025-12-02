package solutions

func CountTrapezoids(points [][]int) int {
	pointNum := make(map[int]int)
	mod := 1000000007
	res, sum := 0, 0

	for _, point := range points {
		y := point[1]
		pointNum[y]++
	}

	for _, pNum := range pointNum {
		edge := pNum * (pNum - 1) / 2
		res = (res + edge*sum) % mod
		sum = (sum + edge) % mod
	}

	return res
}

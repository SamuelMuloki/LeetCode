package solutions

import "math"

func CountTriples(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			sum := (i * i) + (j * j)
			k := int(math.Sqrt(float64(sum)))
			if k <= n && k*k == sum {
				res++
			}
		}
	}

	return res
}

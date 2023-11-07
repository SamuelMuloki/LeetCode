package solutions

import "sort"

func EliminateMaximum(dist []int, speed []int) int {
	arrival := make([]float64, 0)
	for i := 0; i < len(dist); i++ {
		arrival = append(arrival, float64(dist[i])/float64(speed[i]))
	}
	sort.Float64s(arrival)

	ans := 0
	for i := range arrival {
		if arrival[i] <= float64(i) {
			break
		}

		ans++
	}

	return ans
}

package solutions

import "math"

func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cost := make([]int, n)
	for i := range cost {
		cost[i] = math.MaxInt
	}

	cost[src] = 0
	for i := 0; i <= k; i++ {
		temp := append([]int{}, cost...)
		for _, f := range flights {
			curr, next, price := f[0], f[1], f[2]
			if cost[curr] == math.MaxInt {
				continue
			}
			temp[next] = min(temp[next], cost[curr]+price)
		}

		cost = temp
	}

	if cost[dst] == math.MaxInt {
		return -1
	}

	return cost[dst]
}

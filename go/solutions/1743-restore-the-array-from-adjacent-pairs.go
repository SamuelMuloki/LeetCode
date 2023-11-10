package solutions

import "math"

func RestoreArray(adjacentPairs [][]int) []int {
	graph := make(map[int][]int)
	for i := range adjacentPairs {
		graph[adjacentPairs[i][0]] = append(graph[adjacentPairs[i][0]], adjacentPairs[i][1])
		graph[adjacentPairs[i][1]] = append(graph[adjacentPairs[i][1]], adjacentPairs[i][0])
	}

	root := 0
	for val, pair := range graph {
		if len(pair) == 1 {
			root = val
			break
		}
	}

	curr := root
	ans := make([]int, 0)
	ans = append(ans, root)
	prev := math.MaxInt

	for len(ans) < len(graph) {
		for _, neighbor := range graph[curr] {
			if neighbor != prev {
				ans = append(ans, neighbor)
				prev = curr
				curr = neighbor
				break
			}
		}
	}

	return ans
}

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

	ans := make([]int, 0)
	var dfs func(node, prev int)
	dfs = func(node, prev int) {
		ans = append(ans, node)
		for _, neighbor := range graph[node] {
			if neighbor != prev {
				dfs(neighbor, node)
			}
		}
	}

	dfs(root, math.MaxInt)
	return ans
}

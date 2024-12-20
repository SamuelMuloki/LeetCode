package solutions

import (
	"math"
)

const CHAR_COUNT = 26
const INF = math.MaxInt32 / 2

func MinimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	conversionGraph := buildConversionGraph(original, changed, cost)
	optimizeConversionPaths(conversionGraph)
	return computeTotalConversionCost(source, target, conversionGraph)
}

func buildConversionGraph(original []byte, changed []byte, cost []int) [][]int {
	graph := make([][]int, CHAR_COUNT)
	for i := range graph {
		graph[i] = make([]int, CHAR_COUNT)
		for j := range graph[i] {
			graph[i][j] = INF
		}
		graph[i][i] = 0
	}
	for i, c := range cost {
		from := int(original[i] - 'a')
		to := int(changed[i] - 'a')
		if c < graph[from][to] {
			graph[from][to] = c
		}
	}
	return graph
}

func optimizeConversionPaths(graph [][]int) {
	for k := 0; k < CHAR_COUNT; k++ {
		for i := 0; i < CHAR_COUNT; i++ {
			if graph[i][k] < INF {
				for j := 0; j < CHAR_COUNT; j++ {
					if graph[k][j] < INF {
						if graph[i][k]+graph[k][j] < graph[i][j] {
							graph[i][j] = graph[i][k] + graph[k][j]
						}
					}
				}
			}
		}
	}
}

func computeTotalConversionCost(source string, target string, graph [][]int) int64 {
	var totalCost int64 = 0
	for i := 0; i < len(source); i++ {
		sourceChar := int(source[i] - 'a')
		targetChar := int(target[i] - 'a')
		if sourceChar != targetChar {
			if graph[sourceChar][targetChar] == INF {
				return -1
			}
			totalCost += int64(graph[sourceChar][targetChar])
		}
	}
	return totalCost
}

package solutions

import "math"

func ClosestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
	dist := func(start int) []int {
		d := make([]int, n)
		for i := range d {
			d[i] = math.MaxInt32
		}
		steps, curr := 0, start
		for curr != -1 && d[curr] == math.MaxInt32 {
			d[curr] = steps
			curr = edges[curr]
			steps++
		}
		return d
	}

	dist1 := dist(node1)
	dist2 := dist(node2)

	minDistNode, minDist := -1, math.MaxInt32
	for i := 0; i < n; i++ {
		maxDist := dist1[i]
		if dist2[i] > maxDist {
			maxDist = dist2[i]
		}
		if dist1[i] != math.MaxInt32 && dist2[i] != math.MaxInt32 && maxDist < minDist {
			minDist = maxDist
			minDistNode = i
		}
	}
	return minDistNode
}

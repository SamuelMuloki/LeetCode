package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func MinCostClimbingStairs(cost []int) int {
	n := len(cost)
	cost = append(cost, 0)
	for i := n - 3; i >= 0; i-- {
		cost[i] += utils.Min(cost[i+1], cost[i+2])
	}

	return utils.Min(cost[0], cost[1])
}

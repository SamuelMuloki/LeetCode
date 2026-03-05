package solutions

import "math"

func JudgePoint24(cards []int) bool {
	const EPS = 1e-6

	var dfs func([]float64) bool
	dfs = func(nums []float64) bool {
		if len(nums) == 1 {
			return math.Abs(nums[0]-24.0) < EPS
		}

		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums); j++ {
				if i == j {
					continue
				}

				next := []float64{}
				for k := 0; k < len(nums); k++ {
					if k != i && k != j {
						next = append(next, nums[k])
					}
				}

				a, b := nums[i], nums[j]
				candidates := []float64{a + b, a - b, b - a, a * b}
				if math.Abs(b) > EPS {
					candidates = append(candidates, a/b)
				}
				if math.Abs(a) > EPS {
					candidates = append(candidates, b/a)
				}

				for _, val := range candidates {
					if dfs(append(next, val)) {
						return true
					}
				}
			}
		}
		return false
	}

	nums := make([]float64, len(cards))
	for i, v := range cards {
		nums[i] = float64(v)
	}

	return dfs(nums)
}

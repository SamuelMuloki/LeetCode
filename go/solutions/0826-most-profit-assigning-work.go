package solutions

import "sort"

type job struct {
	difficulty int
	profit     int
}

func MaxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	var jobs []job
	for i := range profit {
		jobs = append(jobs, job{difficulty[i], profit[i]})
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].difficulty < jobs[j].difficulty
	})

	sort.Ints(worker)
	var ans, idx, maxProfit int
	for i := range worker {
		for idx < len(jobs) && worker[i] >= jobs[idx].difficulty {
			maxProfit = max(maxProfit, jobs[idx].profit)
			idx++
		}
		ans += maxProfit
	}

	return ans
}

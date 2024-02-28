package solutions

import "sort"

type Job struct {
	start, end, profit int
}

func JobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([]Job, len(startTime))
	for i := 0; i < len(startTime); i++ {
		jobs[i] = Job{
			start:  startTime[i],
			end:    endTime[i],
			profit: profit[i],
		}
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})

	maxProfit := make([]int, len(startTime))
	maxProfit[0] = jobs[0].profit

	for i := 1; i < len(startTime); i++ {
		var prevProfit int
		first := sort.Search(i, func(j int) bool { return jobs[i].start < jobs[j].end })
		if first == 0 {
			prevProfit = 0
		} else {
			prevProfit = maxProfit[first-1]
		}

		maxProfit[i] = prevProfit + jobs[i].profit
		if maxProfit[i-1] > maxProfit[i] {
			maxProfit[i] = maxProfit[i-1]
		}
	}

	return maxProfit[len(maxProfit)-1]
}

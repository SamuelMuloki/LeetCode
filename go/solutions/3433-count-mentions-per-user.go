package solutions

import (
	"sort"
	"strconv"
	"strings"
)

func CountMentions(numberOfUsers int, events [][]string) []int {
	sort.Slice(events, func(i, j int) bool {
		timeA, _ := strconv.Atoi(events[i][1])
		timeB, _ := strconv.Atoi(events[j][1])
		if timeA != timeB {
			return timeA < timeB
		}
		return events[i][0] != "MESSAGE" && events[j][0] == "MESSAGE"
	})

	res := make([]int, numberOfUsers)
	nextOnlineTime := make([]int, numberOfUsers)

	for _, event := range events {
		curr, _ := strconv.Atoi(event[1])

		switch event[0] {
		case "OFFLINE":
			idx, _ := strconv.Atoi(event[2])
			nextOnlineTime[idx] = curr + 60
		case "MESSAGE":
			switch event[2] {
			case "ALL":
				for i := range res {
					res[i]++
				}
			case "HERE":
				for i := range nextOnlineTime {
					if nextOnlineTime[i] <= curr {
						res[i]++
					}
				}
			default:
				users := strings.Split(event[2], " ")
				for _, user := range users {
					num, _ := strconv.Atoi(user[2:])
					res[num]++
				}
			}
		}
	}

	return res
}

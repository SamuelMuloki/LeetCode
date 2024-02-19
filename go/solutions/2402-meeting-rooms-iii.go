package solutions

import (
	"math"
	"sort"
)

func MostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	// rooms[0] - endTime
	// rooms[1] - meetingCounts
	rooms := make([][2]int, n+1)
	rooms[n][0] = math.MaxInt
	result := 0

	for _, meeting := range meetings {
		start := meeting[0]
		end := meeting[1]

		minIndex := n
		flag := true

		for i := 0; i < n; i++ {
			if rooms[i][0] <= start {
				// Assigning meeting to the first vacant room
				rooms[i][0] = end
				rooms[i][1]++
				flag = false
				break
			}
			if rooms[i][0] < rooms[minIndex][0] {
				minIndex = i
			}
		}

		if flag {
			// Time travel to the occupied room with least end time.
			rooms[minIndex][0] = end + (rooms[minIndex][0] - start)
			rooms[minIndex][1]++
		}
	}

	for i := 0; i < n; i++ {
		if rooms[i][1] > rooms[result][1] {
			result = i
		}
	}

	return result
}

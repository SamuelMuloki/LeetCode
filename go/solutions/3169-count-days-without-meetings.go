package solutions

import "sort"

type Event struct {
	day, delta int
}

func CountDays(days int, meetings [][]int) int {
	events := make([]Event, 0, len(meetings)*2)

	for _, meeting := range meetings {
		events = append(events, Event{meeting[0], 1})
		events = append(events, Event{meeting[1] + 1, -1})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].day < events[j].day ||
			(events[i].day == events[j].day && events[i].delta < events[j].delta)
	})

	freeDays := 0
	activeMeetings := 0
	prevDay := 1
	for _, event := range events {
		currDay := event.day
		if activeMeetings == 0 && currDay > prevDay {
			freeDays += currDay - prevDay
		}

		activeMeetings += event.delta
		prevDay = currDay
	}

	if activeMeetings == 0 && prevDay <= days {
		freeDays += days - prevDay + 1
	}

	return freeDays
}

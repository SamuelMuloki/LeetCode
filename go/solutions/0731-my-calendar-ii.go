package solutions

type pair struct {
	st int
	et int
}

type MyCalendarTwo struct {
	events       []pair
	doubleBooked []pair
}

func MyCalendarTwoConstructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (m *MyCalendarTwo) Book(start int, end int) bool {
	for _, val := range m.doubleBooked {
		if val.et > start && val.st < end {
			return false
		}
	}

	for _, val := range m.events {
		if val.et > start && val.st < end {
			m.doubleBooked = append(m.doubleBooked, pair{max(val.st, start), min(val.et, end)})
		}
	}
	m.events = append(m.events, pair{start, end})
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

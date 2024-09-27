package solutions

type MyCalendar struct {
	bookings [][]int
}

func MyCalendarConstructor() MyCalendar {
	return MyCalendar{
		bookings: make([][]int, 0),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, booking := range this.bookings {
		if start < booking[1] && booking[0] < end {
			return false
		}
	}
	this.bookings = append(this.bookings, []int{start, end})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

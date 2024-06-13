package solutions

import "sort"

func MinMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	ans := 0
	for i := 0; i < len(seats); i++ {
		ans += abs(students[i] - seats[i])
	}

	return ans
}

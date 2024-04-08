package solutions

func CountStudents(students []int, sandwiches []int) int {
	count := make(map[int]int)
	for i := 0; i < len(students); i++ {
		count[students[i]]++
	}

	for i := 0; i < len(sandwiches); i++ {
		if count[sandwiches[i]] == 0 {
			return len(sandwiches) - i
		}

		count[sandwiches[i]]--
	}

	return 0
}

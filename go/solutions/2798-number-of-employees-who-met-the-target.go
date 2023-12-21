package solutions

func NumberOfEmployeesWhoMetTarget(hours []int, target int) int {
	ans := 0
	for i := range hours {
		if hours[i] < target {
			continue
		}

		ans++
	}
	return ans
}

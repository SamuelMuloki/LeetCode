package solutions

func FindSpecialInteger(arr []int) int {
	times := len(arr) / 4
	set := make(map[int]int)
	for i := range arr {
		set[arr[i]]++

		if set[arr[i]] > times {
			return arr[i]
		}
	}

	return -1
}

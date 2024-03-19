package solutions

func LeastInterval(tasks []byte, n int) int {
	freq := [26]int{}
	maxCount := 0

	for _, task := range tasks {
		freq[task-'A']++
		maxCount = max(maxCount, freq[task-'A'])
	}

	time := (maxCount - 1) * (n + 1)
	for _, f := range freq {
		if f == maxCount {
			time++
		}
	}

	return max(len(tasks), time)
}

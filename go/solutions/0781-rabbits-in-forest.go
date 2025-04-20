package solutions

import "math"

func NumRabbits(answers []int) int {
	freq := make(map[int]int)
	for _, ans := range answers {
		freq[ans]++
	}

	res := 0
	for num, count := range freq {
		groupSize := num + 1
		numGroups := int(math.Ceil(float64(count) / float64(groupSize)))
		res += numGroups * groupSize
	}

	return res
}

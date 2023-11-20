package solutions

func GarbageCollection(garbage []string, travel []int) int {
	prefixSum := make([]int, len(travel)+1)
	prefixSum[1] = travel[0]
	for i := 1; i < len(travel); i++ {
		prefixSum[i+1] = prefixSum[i] + travel[i]
	}

	garbageLastPos := make(map[byte]int)
	garbageCount := make(map[byte]int)
	for i := range garbage {
		for j := range garbage[i] {
			garbageLastPos[garbage[i][j]] = i
			garbageCount[garbage[i][j]]++
		}
	}

	garbageTypes := []byte{'M', 'P', 'G'}
	ans := 0
	for _, ch := range garbageTypes {
		if count, ok := garbageCount[ch]; ok {
			ans += prefixSum[garbageLastPos[ch]] + count
		}
	}

	return ans
}

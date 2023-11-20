package solutions

func GarbageCollection(garbage []string, travel []int) int {
	for i := 1; i < len(travel); i++ {
		travel[i] = travel[i-1] + travel[i]
	}

	garbageLastPos := make(map[rune]int)
	ans := 0
	for i := range garbage {
		for _, ch := range garbage[i] {
			garbageLastPos[ch] = i
		}

		ans += len(garbage[i])
	}

	garbageTypes := "MPG"
	for _, ch := range garbageTypes {
		if garbageLastPos[ch] != 0 {
			ans += travel[garbageLastPos[ch]-1]
		}
	}

	return ans
}

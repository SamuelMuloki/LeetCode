package solutions

func PoorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	testsPerPig := minutesToTest / minutesToDie
	numPigs, states := 0, 1
	for states < buckets {
		states *= (testsPerPig + 1)
		numPigs++
	}
	return numPigs
}

package solutions

func SmallestDistancePair(nums []int, k int) int {
	arraySize := len(nums)
	maxElement := 0
	for i := range nums {
		maxElement = max(maxElement, nums[i])
	}

	var abs = func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	distanceBucket := make([]int, maxElement+1)
	for i := 0; i < arraySize; i++ {
		for j := i + 1; j < arraySize; j++ {
			distance := abs(nums[i] - nums[j])

			distanceBucket[distance]++
		}
	}

	for dist := 0; dist <= maxElement; dist++ {
		k -= distanceBucket[dist]

		if k <= 0 {
			return dist
		}
	}

	return -1
}

package solutions

import "math"

func MaxDistance2(arrays [][]int) int {
	max1, maxIdx := math.MinInt, 0
	max2, max2Idx := math.MinInt, 0
	for i, arr := range arrays {
		if arr[len(arr)-1] > max1 {
			max2 = max(max2, max1)
			max1 = arr[len(arr)-1]
			maxIdx = i
		} else if arr[len(arr)-1] > max2 {
			max2 = arr[len(arr)-1]
			max2Idx = i
		}
	}

	min1, min2 := math.MaxInt, math.MaxInt
	for i, arr := range arrays {
		if arr[0] < min1 && maxIdx != i {
			min2 = min(min2, min1)
			min1 = arr[0]
		} else if arr[0] < min2 && max2Idx != i {
			min2 = arr[0]
		}
	}

	return max(abs(max1-min1), abs(max2-min2))
}

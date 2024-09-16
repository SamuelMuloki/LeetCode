package solutions

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func FindMinDifference(timePoints []string) int {
	minutes := make([]int, len(timePoints))
	for i, point := range timePoints {
		arr := strings.Split(point, ":")
		hr, _ := strconv.Atoi(arr[0])
		min, _ := strconv.Atoi(arr[1])
		minutes[i] = hr*60 + min
	}
	sort.Ints(minutes)

	ans := math.MaxInt
	for i := 1; i < len(minutes); i++ {
		ans = min(ans, minutes[i]-minutes[i-1])
	}

	return min(ans, 24*60-minutes[len(minutes)-1]+minutes[0])
}

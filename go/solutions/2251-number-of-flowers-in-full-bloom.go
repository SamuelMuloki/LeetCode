package solutions

import "sort"

func FullBloomFlowers(flowers [][]int, people []int) []int {
	starts, ends := make([]int, len(flowers)), make([]int, len(flowers))
	for i, flower := range flowers {
		starts[i] = flower[0]
		ends[i] = flower[1] + 1
	}

	sort.Ints(starts)
	sort.Ints(ends)

	res := make([]int, len(people))
	for k := 0; k < len(people); k++ {
		i := binarySearch(starts, people[k])
		j := binarySearch(ends, people[k])
		res[k] = i - j
	}

	return res
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) / 2
		if target < arr[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

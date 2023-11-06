package solutions

import "sort"

func ArrayRankTransform(arr []int) []int {
	newArr := append([]int{}, arr...)
	sort.Ints(newArr)

	set := make(map[int]int)
	rank := 1
	for i := range newArr {
		if _, ok := set[newArr[i]]; !ok {
			set[newArr[i]] = rank
			rank++
		}
	}

	for i := range arr {
		arr[i] = set[arr[i]]
	}

	return arr
}

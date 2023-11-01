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

	ans := make([]int, len(arr))
	for i := range arr {
		ans[i] = set[arr[i]]
	}

	return ans
}

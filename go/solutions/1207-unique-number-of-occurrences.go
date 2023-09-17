package solutions

func UniqueOccurrences(arr []int) bool {
	unique := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		unique[arr[i]]++
	}

	valSet := make(map[int]int)
	for _, v := range unique {
		valSet[v]++
	}

	return len(valSet) == len(unique)
}

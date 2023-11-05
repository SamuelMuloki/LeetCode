package solutions

func GetWinner(arr []int, k int) int {
	set := make(map[int]int)
	for i := range arr {
		set[arr[i]] = 0
	}

	n := len(arr)
	for i := 1; i < n; i++ {
		arr[1] = arr[i]
		if arr[0] < arr[1] {
			arr[0] = arr[1]
		}

		set[arr[0]]++
		if set[arr[0]] == k {
			return arr[0]
		}
	}

	return arr[0]
}

package solutions

func MaxChunksToSorted(arr []int) int {
	n := len(arr)
	st := []int{}
	for i := 0; i < n; i++ {
		if len(st) == 0 || len(st) > 0 && st[len(st)-1] < arr[i] {
			st = append(st, arr[i])
		} else {
			maxElement := st[len(st)-1]
			for len(st) > 0 && st[len(st)-1] > arr[i] {
				st = st[:len(st)-1]
			}
			st = append(st, maxElement)
		}
	}

	return len(st)
}

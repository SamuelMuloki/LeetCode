package solutions

func FindSpecialInteger(arr []int) int {
	size := len(arr) / 4
	for i := 0; i < len(arr)-size; i++ {
		if arr[i] == arr[i+size] {
			return arr[i]
		}
	}

	return -1
}

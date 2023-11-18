package solutions

func ReplaceElements(arr []int) []int {
	prev := -1
	for i := len(arr) - 1; i >= 0; i-- {
		curr := arr[i]
		arr[i] = prev
		prev = max(prev, curr)
	}

	return arr
}

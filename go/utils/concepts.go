package utils

func MaxSum(arr []int, k int) int {
	max_sum := 0

	for i := 0; i < k; i++ {
		max_sum += arr[i]
	}

	window_sum := max_sum
	for j := k; j < len(arr); j++ {
		window_sum += arr[j] - arr[j-k]
		max_sum = Max(window_sum, max_sum)
	}

	return max_sum
}

// Reverse a string

func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func Min(x int, y int) int {
	if x < y {
		return x
	}

	return y
}

func Max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

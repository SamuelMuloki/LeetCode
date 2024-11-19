package solutions

func Construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}

	resultArray := make([][]int, m)
	for i := range resultArray {
		resultArray[i] = make([]int, n)
	}

	for i := 0; i < len(original); i++ {
		resultArray[i/n][i%n] = original[i]
	}

	return resultArray
}

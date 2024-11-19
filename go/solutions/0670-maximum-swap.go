package solutions

import "strconv"

func MaximumSwap(num int) int {
	arr := []byte(strconv.Itoa(num))
	n := len(arr)
	last := make([]int, 10)

	for i := 0; i < n; i++ {
		last[arr[i]-'0'] = i
	}

	for i := 0; i < n-1; i++ {
		for j := 9; j > int(arr[i]-'0'); j-- {
			if last[j] > i {
				arr[i], arr[last[j]] = arr[last[j]], arr[i]
				res, _ := strconv.Atoi(string(arr))
				return res
			}
		}
	}

	return num
}

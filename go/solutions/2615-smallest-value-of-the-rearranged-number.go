package solutions

import (
	"sort"
	"strconv"
)

func SmallestNumber(num int64) int64 {
	arr := []byte(strconv.Itoa(abs(int(num))))
	sort.Slice(arr, func(i, j int) bool {
		if num < 0 {
			return arr[i] > arr[j]
		}

		return arr[i] < arr[j]
	})

	idx := sort.Search(len(arr), func(i int) bool { return arr[i] != '0' })
	if idx != len(arr) {
		arr[0], arr[idx] = arr[idx], arr[0]
	}

	ans, _ := strconv.Atoi(string(arr))
	if num < 0 {
		return int64(-ans)
	}

	return int64(ans)
}

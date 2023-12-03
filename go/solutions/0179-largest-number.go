package solutions

import (
	"sort"
	"strconv"
	"strings"
)

func LargestNumber(nums []int) string {
	arr := make([]string, len(nums))
	for i := range nums {
		arr[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i]+arr[j] > arr[j]+arr[i]
	})

	if arr[0] == "0" {
		return "0"
	}

	return strings.Join(arr, "")
}

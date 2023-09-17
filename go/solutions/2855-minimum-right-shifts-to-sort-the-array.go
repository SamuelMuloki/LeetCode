package solutions

import (
	"reflect"
	"sort"
)

func MinimumRightShifts(nums []int) int {
	l, r := 0, len(nums)-1

	newArr := append([]int{}, nums...)
	sort.Ints(newArr)

	if reflect.DeepEqual(nums, newArr) {
		return 0
	}

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	arr := append(nums[l:], nums[:l]...)
	if reflect.DeepEqual(arr, newArr) {
		return len(nums) - l
	}

	return -1
}

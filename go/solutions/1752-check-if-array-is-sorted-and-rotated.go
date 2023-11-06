package solutions

import (
	"reflect"
	"sort"
)

func Check(nums []int) bool {
	l, r := 0, len(nums)-1
	newArr := append([]int{}, nums...)
	sort.Ints(newArr)

	if reflect.DeepEqual(nums, newArr) {
		return true
	}

	for l < r && nums[l] == nums[l+1] {
		l++
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
	return reflect.DeepEqual(arr, newArr)
}

package solutions

import "math"

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

type MountainArray struct {
	arr []int
}

func NewMountainArray(arr []int) *MountainArray {
	return &MountainArray{
		arr: arr,
	}
}

func (this *MountainArray) get(index int) int {
	return this.arr[index]
}

func (this *MountainArray) length() int {
	return len(this.arr)
}

func FindInMountainArray(target int, mountainArr *MountainArray) int {
	arrLen := mountainArr.length() - 1
	cache := make(map[int]int)

	peak := findPeak(0, arrLen, mountainArr, cache)

	left := binarySearchLeft(0, peak, target, mountainArr, cache)
	if left != -1 {
		return left
	}

	right := binarySearchRight(peak, arrLen, target, mountainArr, cache)
	if right != -1 {
		return right
	}

	return -1
}

func findPeak(l, r int, mountainArr *MountainArray, cache map[int]int) int {
	for l < r {
		mid := l + (r-l)/2

		var midVal, nextMidVal int
		if val, found := cache[mid]; found {
			midVal = val
		} else {
			midVal = mountainArr.get(mid)
			cache[mid] = midVal
		}

		if val, found := cache[mid+1]; found {
			nextMidVal = val
		} else {
			nextMidVal = mountainArr.get(mid + 1)
			cache[mid+1] = nextMidVal
		}

		if midVal > nextMidVal {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func binarySearchLeft(l, r, target int, mountainArr *MountainArray, cache map[int]int) int {
	for l != r {
		mid := int(math.Ceil(float64((l + r)) / 2))

		var midVal int
		if val, found := cache[mid]; found {
			midVal = val
		} else {
			midVal = mountainArr.get(mid)
			cache[mid] = midVal
		}

		if midVal > target {
			r = mid - 1
		} else {
			l = mid
		}
	}

	if mountainArr.get(l) == target {
		return l
	}

	return -1
}

func binarySearchRight(l, r, target int, mountainArr *MountainArray, cache map[int]int) int {
	for l != r {
		mid := int(math.Ceil(float64((l + r)) / 2))

		var midVal int
		if val, found := cache[mid]; found {
			midVal = val
		} else {
			midVal = mountainArr.get(mid)
			cache[mid] = midVal
		}

		if midVal < target {
			r = mid - 1
		} else {
			l = mid
		}
	}

	var targetVal int
	if val, found := cache[l]; found {
		targetVal = val
	} else {
		targetVal = mountainArr.get(l)
		cache[l] = targetVal
	}

	if targetVal == target {
		return l
	}

	return -1
}

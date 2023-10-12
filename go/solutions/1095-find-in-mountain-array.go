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
	peak := findPeak(0, arrLen, mountainArr)

	left := binarySearchLeft(0, peak, target, mountainArr)
	if left != -1 {
		return left
	}

	right := binarySearchRight(peak, arrLen, target, mountainArr)
	if right != -1 {
		return right
	}

	return -1
}

func findPeak(l, r int, mountainArr *MountainArray) int {
	for l < r {
		mid := l + (r-l)/2
		if mountainArr.get(mid) > mountainArr.get(mid+1) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func binarySearchLeft(l, r, target int, mountainArr *MountainArray) int {
	for l != r {
		mid := int(math.Ceil(float64((l + r)) / 2))
		if mountainArr.get(mid) > target {
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

func binarySearchRight(l, r, target int, mountainArr *MountainArray) int {
	for l != r {
		mid := int(math.Ceil(float64((l + r)) / 2))
		if mountainArr.get(mid) < target {
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

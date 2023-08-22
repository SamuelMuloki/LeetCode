package solutions

import (
	"math"

	"github.com/SamuelMuloki/LeetCode/go/utils"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2

	if len(A) > len(B) {
		A, B = B, A
	}
	total := len(A) + len(B)
	mid := (total + 1) / 2

	l, r := 0, len(A)-1
	for {
		i := (l + r) >> 1
		j := mid - i - 2

		var ALeft, ARight, BLeft, BRight int
		if i >= 0 {
			ALeft = A[i]
		} else {
			ALeft = -math.MaxInt
		}
		if j >= 0 {
			ARight = B[j]
		} else {
			ARight = -math.MaxInt
		}
		if i+1 < len(A) {
			BLeft = A[i+1]
		} else {
			BLeft = math.MaxInt
		}
		if j+1 < len(B) {
			BRight = B[j+1]
		} else {
			BRight = math.MaxInt
		}

		if ALeft <= BRight && ARight <= BLeft {
			if total%2 == 0 {
				return float64(utils.Max(ALeft, ARight)+utils.Min(BLeft, BRight)) / 2
			}
			return float64(utils.Max(ALeft, ARight))
		}

		if ALeft > BRight {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}

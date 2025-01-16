package solutions

func XorAllNums(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	num1Xor := 0
	if n2&1 == 1 {
		for j := 0; j < len(nums1); j++ {
			num1Xor ^= nums1[j]
		}
	}

	num2Xor := 0
	if n1&1 == 1 {
		for j := 0; j < len(nums2); j++ {
			num2Xor ^= nums2[j]
		}
	}

	return num1Xor ^ num2Xor
}

package solutions

func MinSum(nums1 []int, nums2 []int) int64 {
	var s1, s2, z1, z2 int64
	for i := 0; i < len(nums1); i++ {
		s1 += int64(nums1[i])
		if nums1[i] == 0 {
			z1++
		}
	}

	for i := 0; i < len(nums2); i++ {
		s2 += int64(nums2[i])
		if nums2[i] == 0 {
			z2++
		}
	}

	if s1+z1 < s2+z2 && z1 > 0 {
		return s2 + z2
	} else if s1+z1 > s2+z2 && z2 > 0 {
		return s1 + z1
	} else if s1+z1 == s2+z2 {
		return s1 + z1
	}

	return -1
}

package solutions

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}

	for i, num := range nums1 {
		found := false
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == num {
				found = true
			}

			if found && nums2[j] > num {
				ans[i] = nums2[j]
				break
			}
		}
	}

	return ans
}

package solutions

func FindDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		set1[nums1[i]]++
	}

	set2 := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		set2[nums2[i]]++
	}

	arr1 := []int{}
	for key := range set1 {
		if set2[key] == 0 {
			arr1 = append(arr1, key)
		}
	}

	arr2 := []int{}
	for key := range set2 {
		if set1[key] == 0 {
			arr2 = append(arr2, key)
		}
	}

	return [][]int{arr1, arr2}
}

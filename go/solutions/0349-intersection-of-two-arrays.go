package solutions

func Intersection(nums1 []int, nums2 []int) []int {
	output := make([]int, 0)
	set := make(map[int]bool)

	for _, num := range nums1 {
		set[num] = true
	}

	for _, num := range nums2 {
		if set[num] {
			output = append(output, num)
			set[num] = false
		}
	}

	return output
}

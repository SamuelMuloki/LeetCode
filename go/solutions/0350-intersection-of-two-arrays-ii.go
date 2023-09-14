package solutions

func Intersect(nums1 []int, nums2 []int) []int {
	output := make([]int, 0)
	set := make(map[int]int)

	minArr, maxArr := nums1, nums2
	if len(nums1) > len(nums2) {
		minArr, maxArr = nums2, nums1
	}

	for _, num := range minArr {
		set[num]++
	}

	for _, num := range maxArr {
		if set[num] != 0 {
			output = append(output, num)
			set[num]--
		}
	}

	return output
}

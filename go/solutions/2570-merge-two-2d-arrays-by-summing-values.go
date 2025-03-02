package solutions

func MergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	set := [1001]int{}
	for _, num := range nums1 {
		set[num[0]] += num[1]
	}

	for _, num := range nums2 {
		set[num[0]] += num[1]
	}

	ans := make([][]int, 0)
	for i, num := range set {
		if num == 0 {
			continue
		}
		ans = append(ans, []int{i, num})
	}

	return ans
}

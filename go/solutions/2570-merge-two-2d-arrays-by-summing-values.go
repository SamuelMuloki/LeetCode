package solutions

func MergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	set := [1001]int{}
	maxCnt := max(len(nums1), len(nums2))
	for i := 0; i < maxCnt; i++ {
		if i < len(nums1) {
			set[nums1[i][0]] += nums1[i][1]
		}
		if i < len(nums2) {
			set[nums2[i][0]] += nums2[i][1]
		}
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

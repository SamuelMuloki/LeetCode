package solutions

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	ng, st := map[int]int{}, []int{}
	for _, num := range nums2 {
		ng[num] = -1
		for len(st) > 0 && st[len(st)-1] < num {
			top := st[len(st)-1]
			st = st[:len(st)-1]
			ng[top] = num
		}
		st = append(st, num)
	}

	ans := []int{}
	for _, num := range nums1 {
		ans = append(ans, ng[num])
	}

	return ans
}

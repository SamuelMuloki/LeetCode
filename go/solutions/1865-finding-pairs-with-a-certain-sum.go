package solutions

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	set   map[int]int
}

func FindSumPairsConstructor(nums1 []int, nums2 []int) FindSumPairs {
	set := make(map[int]int)
	for _, num2 := range nums2 {
		set[num2]++
	}
	return FindSumPairs{
		nums1,
		nums2,
		set,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	curr := this.nums2[index]
	this.set[curr]--
	this.nums2[index] += val
	this.set[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	cnt := 0
	for _, num1 := range this.nums1 {
		if val, ok := this.set[tot-num1]; ok {
			cnt += val
		}
	}

	return cnt
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */

package solutions

type BIT struct {
	tree []int
	n    int
}

func NewBIT(n int) *BIT {
	return &BIT{
		tree: make([]int, n+1),
		n:    n,
	}
}

func (bit *BIT) Update(index, delta int) {
	index++
	for index <= bit.n {
		bit.tree[index] += delta
		index += index & -index
	}
}

func (bit *BIT) Query(index int) int {
	index++
	sum := 0
	for index > 0 {
		sum += bit.tree[index]
		index -= index & -index
	}
	return sum
}

func GoodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	valueToIndex := make(map[int]int)
	for i, val := range nums2 {
		valueToIndex[val] = i
	}

	p := make([]int, n)
	for i, val := range nums1 {
		p[i] = valueToIndex[val]
	}

	leftBIT := NewBIT(n)
	rightBIT := NewBIT(n)
	for _, idx := range p {
		rightBIT.Update(idx, 1)
	}

	var res int64
	for j := 0; j < n; j++ {
		pj := p[j]
		rightBIT.Update(pj, -1)
		leftCount := leftBIT.Query(pj - 1)
		rightCount := rightBIT.Query(n-1) - rightBIT.Query(pj)
		res += int64(leftCount) * int64(rightCount)
		leftBIT.Update(pj, 1)
	}
	return res
}

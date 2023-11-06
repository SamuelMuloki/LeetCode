package solutions

func SortArray(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, l, r int) {
	if l < r {
		mid := l + (r-l)/2

		mergeSort(nums, l, mid)
		mergeSort(nums, mid+1, r)

		merge(nums, l, mid, r)
	}
}

func merge(nums []int, l, m, r int) {
	n1, n2 := m-l+1, r-m

	L, R := make([]int, n1), make([]int, n2)
	for i := 0; i < n1; i++ {
		L[i] = nums[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = nums[m+1+j]
	}

	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			nums[k] = L[i]
			i++
		} else {
			nums[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		nums[k] = L[i]
		i++
		k++
	}

	for i < n2 {
		nums[k] = R[i]
		j++
		k++
	}
}

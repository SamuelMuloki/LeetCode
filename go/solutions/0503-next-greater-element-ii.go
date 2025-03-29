package solutions

func NextGreaterElementsII(nums []int) []int {
	n := len(nums)
	ng := make(map[int]int)
	st := []int{}
	for i := 0; i < n*2; i++ {
		for len(st) > 0 && nums[st[len(st)-1]%n] < nums[i%n] {
			top := st[len(st)-1]
			st = st[:len(st)-1]
			ng[top%n] = i % n
		}
		st = append(st, i)
	}

	arr := make([]int, n)
	for i := range nums {
		if _, ok := ng[i]; ok {
			arr[i] = nums[ng[i]]
		} else {
			arr[i] = -1
		}
	}

	return arr
}

package solutions

func DivideArray2(nums []int) bool {
	set := make(map[int]int)
	cnt := 0
	for _, num := range nums {
		set[num]++
		if set[num]%2 == 0 {
			cnt++
		}
	}

	return cnt == len(nums)/2
}

package solutions

func FindMatrix(nums []int) [][]int {
	set := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		set[nums[i]]++
	}

	ans := make([][]int, 0)
	for len(set) > 0 {
		row := []int{}
		for num := range set {
			row = append(row, num)
			set[num]--
			if set[num] == 0 {
				delete(set, num)
			}
		}
		ans = append(ans, row)
	}

	return ans
}

package solutions

func GetSneakyNumbers(nums []int) []int {
	set := [101]int{}
	res := make([]int, 0, 2)
	for _, num := range nums {
		set[num]++
		if set[num] > 1 {
			res = append(res, num)
		}
		if len(res) == 2 {
			return res
		}
	}

	return []int{}
}

package solutions

func PrefixesDivBy5(nums []int) []bool {
	n := len(nums)
	res := make([]bool, n)

	prefix := 0
	for i, bit := range nums {
		prefix = (prefix*2 + bit) % 5
		res[i] = prefix == 0
	}

	return res
}

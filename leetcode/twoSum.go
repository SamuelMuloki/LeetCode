package leetcode

func TwoSum(nums []int, target int) []int {
	complements := make(map[int]int, len(nums))

	for i, val := range nums {
		if comp, ok := complements[val]; ok {
			return []int{comp, i}
		}

		complements[target-val] = i
	}

	return []int{}
}

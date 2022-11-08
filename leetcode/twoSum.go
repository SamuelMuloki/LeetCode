package leetcode

func TwoSum(nums []int, target int) []int {
	complements := make(map[int]int, len(nums))
	indices := make([]int, 0, 2)

	for i, val := range nums {

		if comp, ok := complements[val]; ok {
			indices = []int{comp, i}
		}

		complements[target-val] = i
	}

	return indices
}

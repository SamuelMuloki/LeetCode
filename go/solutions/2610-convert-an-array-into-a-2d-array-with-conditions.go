package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

func FindMatrix(nums []int) [][]int {
	cnt, maxCnt := [201]int{}, 0
	for i := range nums {
		cnt[nums[i]]++
		maxCnt = utils.Max(maxCnt, cnt[nums[i]])
	}

	ans := make([][]int, maxCnt)
	for i := 0; i < 201; i++ {
		for j := 0; j < cnt[i]; j++ {
			ans[j] = append(ans[j], i)
		}
	}

	return ans
}

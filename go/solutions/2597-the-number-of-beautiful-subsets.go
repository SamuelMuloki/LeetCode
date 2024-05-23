package solutions

import "sort"

func BeautifulSubsets(nums []int, k int) int {
	freqMap := make(map[int]int)
	sort.Ints(nums)
	var countBeautifulSubsets func(i int) int
	countBeautifulSubsets = func(i int) int {
		if i == len(nums) {
			return 1
		}

		count := countBeautifulSubsets(i + 1)
		if _, ok := freqMap[nums[i]-k]; !ok {
			freqMap[nums[i]]++
			count += countBeautifulSubsets(i + 1)
			freqMap[nums[i]]--
			if freqMap[nums[i]] == 0 {
				delete(freqMap, nums[i])
			}
		}

		return count
	}

	return countBeautifulSubsets(0) - 1
}

package solutions

import "sort"

func FindXSum(nums []int, k int, x int) []int {
	res := []int{}
	set := make(map[int]int)
	type kv struct {
		key   int
		value int
	}

	start := 0
	for i, num := range nums {
		set[num]++
		if i >= k-1 {
			var ss []kv
			for k, v := range set {
				ss = append(ss, kv{k, v})
			}

			sort.Slice(ss, func(i, j int) bool {
				if ss[i].value == ss[j].value {
					return ss[i].key > ss[j].key
				}
				return ss[i].value > ss[j].value
			})

			sum := 0
			for i := 0; i < len(ss); i++ {
				if i > x-1 {
					break
				}
				sum += ss[i].key * ss[i].value
			}

			res = append(res, sum)
			set[nums[start]]--
			start++
		}
	}

	return res
}

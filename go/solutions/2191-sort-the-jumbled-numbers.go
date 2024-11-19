package solutions

import "sort"

type Elem struct {
	num, idx int
}

func SortJumbled(mapping []int, nums []int) []int {
	mappedElems := make([]Elem, len(nums))
	for i, num := range nums {
		mappedNum := 0
		if num == 0 {
			mappedNum = mapping[0]
		} else {
			mul := 1
			for ; num > 0; num /= 10 {
				digit := num % 10
				mappedDigit := mapping[digit]
				mappedNum += mappedDigit * mul
				mul *= 10
			}
		}
		mappedElems[i] = Elem{num: mappedNum, idx: i}
	}

	sort.SliceStable(mappedElems, func(i, j int) bool {
		if mappedElems[i].num == mappedElems[j].num {
			return mappedElems[i].idx < mappedElems[j].idx
		}
		return mappedElems[i].num < mappedElems[j].num
	})

	ans := make([]int, len(nums))
	for i, elem := range mappedElems {
		ans[i] = nums[elem.idx]
	}

	return ans
}

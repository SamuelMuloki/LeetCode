package solutions

import "sort"

func FindEvenNumbers(digits []int) []int {
	sort.Ints(digits)
	set := [10]int{}
	for _, digit := range digits {
		set[digit]++
	}

	n := len(digits)
	set2 := make(map[int]bool)
	var res []int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if digits[i] == 0 || i == j || i == k || k == j ||
					digits[k]%2 != 0 {
					continue
				}
				num := digits[i]*100 + digits[j]*10 + digits[k]
				if set2[num] {
					continue
				}
				res = append(res, num)
				set2[num] = true
			}
		}
	}

	return res
}

package solutions

func MajorityElement2(nums []int) []int {
	candidate1, candidate2, count1, count2 := 0, 1, 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 == 0 {
			candidate1 = num
			count1 = 1
		} else if count2 == 0 {
			candidate2 = num
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	var res []int
	if count1 > (len(nums) / 3) {
		res = append(res, candidate1)
	}

	if count2 > (len(nums) / 3) {
		res = append(res, candidate2)
	}

	return res
}

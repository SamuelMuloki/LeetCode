package solutions

import "github.com/SamuelMuloki/LeetCode/go/utils"

type Stream struct {
	Set      map[int]int
	MaxValue int
}

func (s *Stream) GetHighestNumber(num int) int {
	tempMax := 1
	s.Set[num] = num

	sum := 0
	for _, val := range s.Set {
		sum += val
		tempMax = utils.Max(tempMax, val)
	}

	maxSum := getSum(tempMax)
	if maxSum == sum {
		s.MaxValue = tempMax
	}

	return s.MaxValue
}

func getSum(n int) int {
	return n * (n + 1) / 2
}

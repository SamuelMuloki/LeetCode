package leetcode

/*
Given stream of numbers, return the highest number, 1 is the starting point
1,2
GetHighestNumber() -> 2
3, 4
GetHighestNumber() -> 4
5, 7
GetHighestNumber() -> 4
*/

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
		tempMax = max(tempMax, val)
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

package solutions

func MinDominoRotations(tops []int, bottoms []int) int {
	cnt := [7]int{}
	topCnt := [7]int{}
	bottomCnt := [7]int{}

	for i := 0; i < len(tops); i++ {
		topCnt[tops[i]]++
		bottomCnt[bottoms[i]]++
		cnt[tops[i]]++
		if tops[i] != bottoms[i] {
			cnt[bottoms[i]]++
		}
	}

	for i := 1; i <= 6; i++ {
		if cnt[i] == len(tops) {
			return min(len(tops)-topCnt[i], len(tops)-bottomCnt[i])
		}
	}

	return -1
}

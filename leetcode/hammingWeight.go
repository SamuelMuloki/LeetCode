package leetcode

func HammingWeight(num uint32) (count int) {
	for ; num != 0; num, count = num&(num-1), count+1 {
	}
	return
}

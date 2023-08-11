package leetcode

func HammingWeight(num uint32) (count int) {
	for ; num != 0; num, count = num>>1, count+int(num&1) {
	}
	return
}

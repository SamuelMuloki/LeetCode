// https://leetcode.com/problems/power-of-two/
package solutions

func IsPowerOfTwo(n int) bool {
	return (n != 0 && (n&(n-1)) == 0)
}

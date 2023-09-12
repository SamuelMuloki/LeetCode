// https://leetcode.com/problems/guess-number-higher-or-lower
package solutions

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func GuessNumber(n int) int {
	l, r := 0, n

	for {
		mid := l + (r-l)/2
		p := guess(mid)

		if p == -1 {
			r = mid - 1
		} else if p == 1 {
			l = mid + 1
		} else {
			return mid
		}
	}
}

func guess(num int) int {
	pick := 6
	if num > pick {
		return -1
	} else if num < pick {
		return 1
	}

	return 0
}

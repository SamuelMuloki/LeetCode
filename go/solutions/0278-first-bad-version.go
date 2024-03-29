package solutions

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func FirstBadVersion(n int) int {
	l, r := 0, n-1

	for l < r {
		mid := l + (r-l)/2

		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func isBadVersion(version int) bool {
	return version == 4 || version == 5
}

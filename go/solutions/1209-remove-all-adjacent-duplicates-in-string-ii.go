package solutions

import "bytes"

type Pair struct {
	ch   byte
	freq int
}

func RemoveDuplicates4(s string, k int) string {
	st := make([]Pair, 0)
	for i := range s {
		if len(st) > 0 && st[len(st)-1].ch == s[i] {
			st[len(st)-1].freq++
			if st[len(st)-1].freq == k {
				st = st[:len(st)-1]
			}
		} else {
			st = append(st, Pair{s[i], 1})
		}
	}

	var ans string
	for i := range st {
		ans += string(bytes.Repeat([]byte{st[i].ch}, st[i].freq))
	}

	return ans
}

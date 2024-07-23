package solutions

func IsIsomorphic(s string, t string) bool {
	sMap, tMap := [128]byte{}, [128]byte{}
	for i := range s {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		}
		sMap[s[i]], tMap[t[i]] = byte(i+1), byte(i+1)
	}

	return true
}

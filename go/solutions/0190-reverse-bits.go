package solutions

func ReverseBits(num uint32) uint32 {
	mask, res := uint32(1<<31), uint32(0)
	for i := 0; i < 32; i++ {
		if (num & 1) != 0 {
			res |= mask
		}

		mask >>= 1
		num >>= 1
	}

	return res
}

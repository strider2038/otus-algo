package bitboard

func CountBitsSequentially(bits uint64) int {
	count := 0

	for b := bits; b > 0; b = b >> 1 {
		if b&1 == 1 {
			count++
		}
	}

	return count
}

func CountBitsByDivision(bits uint64) int {
	count := 0

	for b := bits; b > 0; b &= b - 1 {
		count++
	}

	return count
}

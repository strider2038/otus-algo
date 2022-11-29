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

func CountBitsByCache(bits uint64) int {
	count := 0

	for i := 0; i < 8 && bits > 0; i++ {
		count += counts8[bits&0xFF]
		bits = bits >> 8
	}

	return count
}

// Предварительно рассчитанный массив счетчиков бит для 8-разрядных чисел
var counts8 = calcCountsByBits8()

func calcCountsByBits8() [0x100]int {
	counts := [0x100]int{}

	for i := 0; i < 0x100; i++ {
		counts[i] = CountBitsByDivision(uint64(i))
	}

	return counts
}

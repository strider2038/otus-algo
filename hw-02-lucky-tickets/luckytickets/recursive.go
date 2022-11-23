package luckytickets

// CountRecursively - универсальный рекурсивный алгоритм для произвольного N.
//
// Сложность O(100 ^ N).
func CountRecursively(N int) int {
	return countRecursively(N, 0, 0, 0)
}

func countRecursively(remainingN, sumA, sumB, count int) int {
	if remainingN == 0 {
		if sumA == sumB {
			count++
		}

		return count
	}

	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			count = countRecursively(remainingN-1, sumA+a, sumB+b, count)
		}
	}

	return count
}

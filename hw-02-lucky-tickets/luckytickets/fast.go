package luckytickets

// CountFast - быстрый алгоритм для произвольного N.
// Основан на алгоритме https://habr.com/ru/post/266479/.
//
// Сложность O(N^2+N) = O(N^2).
func CountFast(N int) int {
	digitsSum := countDigitsSums(N) // O(N^2)

	sumOfSquares := 0
	for i := 0; i < len(digitsSum); i++ { // O(9*N+1) = O(N)
		sumOfSquares += digitsSum[i] * digitsSum[i]
	}

	return sumOfSquares
}

func countDigitsSums(N int) []int {
	if N <= 0 {
		return []int{1}
	}

	maxN := 9*N + 1
	digitsSums := make([]int, maxN)

	previousSums := countDigitsSums(N - 1) // O(N^2)
	for i := 0; i < maxN; i++ {            // O((9*N+1)*10) = O(N)
		k := i
		for j := 0; j < 10 && k >= 0; j++ { // O(10)
			if k < len(previousSums) {
				digitsSums[i] += previousSums[k]
			}
			k--
		}
	}

	return digitsSums
}

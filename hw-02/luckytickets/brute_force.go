package luckytickets

// CountByBruteForceForN3 - подсчитывает количество счастливых билетов для 6-значного числа.
// Наивный метод перебора.
//
// Сложность O(10^N), где N - количество цифр для числа.
func CountByBruteForceForN3() int {
	const N = 10
	count := 0

	for a1 := 0; a1 < N; a1++ {
		for a2 := 0; a2 < N; a2++ {
			for a3 := 0; a3 < N; a3++ {
				for b1 := 0; b1 < N; b1++ {
					for b2 := 0; b2 < N; b2++ {
						for b3 := 0; b3 < N; b3++ {
							sum1 := a1 + a2 + a3
							sum2 := b1 + b2 + b3
							if sum1 == sum2 {
								count++
							}
						}
					}
				}
			}
		}
	}

	return count
}

// CountByBruteForce2ForN3 - подсчитывает количество счастливых билетов для 6-значного числа.
// Альтернативный метод перебора.
//
// Сложность O(10^N), где N - количество цифр для числа.
func CountByBruteForce2ForN3() int {
	const N = 10

	count := 0
	total := N * N * N * N * N * N
	for i := 0; i < total; i++ {
		sum1 := i/100000 + i/10000%10 + i/1000%10
		sum2 := i/100%10 + i/10%10 + i%10
		if sum1 == sum2 {
			count++
		}
	}

	return count
}

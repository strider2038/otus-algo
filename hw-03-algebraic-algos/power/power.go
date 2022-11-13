package power

// Iterative - итеративный алгоритм возведения числа в степень.
// Сложность O(N).
func Iterative(a float64, n int) float64 {
	if n == 0 {
		return 1
	}

	pow := a
	for i := 0; i < n-1; i++ {
		pow *= a
	}

	return pow
}

// Logarithmic - алгоритм возведения в степень через двоичное разложение показателя степени.
// Сложность O(2LogN) = O(LogN).
func Logarithmic(a float64, n int) float64 {
	d := a
	pow := 1.0

	for n > 0 {
		if n&1 == 1 { // эквивалент n % 2 == 1
			pow = pow * d
		}
		n = n >> 1 // эквивалент n = n / 2
		d = d * d
	}

	return pow
}

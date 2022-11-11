package power

func Iterative(a float64, n int) float64 {
	pow := a
	for i := 0; i < n; i++ {
		pow *= a
	}
	return pow
}

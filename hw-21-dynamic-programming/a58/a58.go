package a58

func Calculate(n int) int {
	if n == 0 {
		return 0
	}

	a := []int{
		1, // x5
		0, // x55
		1, // x8
		0, // x88
	}

	for i := 1; i < n; i++ {
		b := []int{
			a[2] + a[3],
			a[0],
			a[0] + a[1],
			a[2],
		}
		a = b
	}

	return sum(a)
}

func sum(n []int) int {
	s := 0
	for _, m := range n {
		s += m
	}
	return s
}

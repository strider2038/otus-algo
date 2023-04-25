package tree

func Calculate(A [][]int) int {
	for i := len(A) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			A[i][j] += max(A[i+1][j], A[i+1][j+1])
		}
	}

	return A[0][0]
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

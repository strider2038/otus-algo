package islands

func Count(M [][]int) int {
	count := 0

	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[i]); j++ {
			if M[i][j] != 0 {
				count++
				setZero(M, i, j)
			}
		}
	}

	return count
}

func setZero(M [][]int, i, j int) {
	if i < 0 || i >= len(M) {
		return
	}
	if j < 0 || j >= len(M[i]) {
		return
	}
	if M[i][j] == 0 {
		return
	}
	M[i][j] = 0
	setZero(M, i-1, j)
	setZero(M, i+1, j)
	setZero(M, i, j-1)
	setZero(M, i, j+1)
}

package pattern

func FindFull(text string, pattern string) int {
	last := len(pattern) - 1

	for i := 0; i < len(text)-last; i++ {
		match := true
		for j := 0; j <= last; j++ {
			match = match && text[i+j] == pattern[j]
		}
		if match {
			return i
		}
	}

	return -1
}

func FindByPatternPrefix(text string, pattern string) int {
	last := len(pattern) - 1

	for i := 0; i < len(text)-last; i++ {
		j := 0
		for ; j <= last && text[i+j] == pattern[j]; j++ {
		}
		if j == len(pattern) {
			return i
		}
	}

	return -1
}

func FindByTextSuffix(text string, pattern string) int {
	last := len(pattern) - 1

	for i := 0; i < len(text)-last; i++ {
		j := last
		for ; j >= 0 && text[i+j] == pattern[j]; j-- {
		}
		if j == -1 {
			return i
		}
	}

	return -1
}

func FindBMH(text string, pattern string) int {
	return FindBMHWithTable(text, pattern, CreateShiftTable(pattern))
}

func FindBMHWithTable(text string, pattern string, shifts [128]int) int {
	last := len(pattern) - 1

	for i := 0; i < len(text)-last; i += shifts[text[i+last]] {
		j := last
		for ; j >= 0 && text[i+j] == pattern[j]; j-- {
		}
		if j == -1 {
			return i
		}
	}

	return -1
}

func CreateShiftTable(pattern string) [128]int {
	shifts := [128]int{}

	for i := 0; i < len(shifts); i++ {
		shifts[i] = len(pattern)
	}
	for i := 0; i < len(pattern); i++ {
		shifts[pattern[i]] = len(pattern) - i - 1
	}

	return shifts
}

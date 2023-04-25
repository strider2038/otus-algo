package goroh

func GCD(a, b int) int {
	if a == b {
		return a
	}
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	if a&1 == 0 && b&1 == 0 {
		return GCD(a>>1, b>>1) << 1
	}
	if a&1 == 0 && b&1 != 0 {
		return GCD(a>>1, b)
	}
	if a&1 != 0 && b&1 == 0 {
		return GCD(a, b>>1)
	}

	if a > b {
		return GCD((a-b)>>1, b)
	}

	return GCD(a, (b-a)>>1)
}

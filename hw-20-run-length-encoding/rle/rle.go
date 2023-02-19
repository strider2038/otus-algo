package rle

func CompressBytes(input []byte) []byte {
	if len(input) == 0 {
		return input
	}

	output := make([]byte, 0, len(input))

	cursor := input[0]
	count := byte(1)
	for i := 1; i < len(input); i++ {
		if input[i] == cursor && count < 255 {
			count++
		} else {
			output = append(output, count, cursor)
			cursor = input[i]
			count = 1
		}
	}

	output = append(output, count, cursor)

	return output
}

func DecompressBytes(input []byte) []byte {
	if len(input) == 0 {
		return input
	}
	if len(input)&1 != 0 {
		panic("unexpected length of input")
	}

	output := make([]byte, 0, len(input))
	for i := 0; i < len(input); i += 2 {
		count := input[i]
		cursor := input[i+1]
		for j := 0; j < int(count); j++ {
			output = append(output, cursor)
		}
	}

	return output
}

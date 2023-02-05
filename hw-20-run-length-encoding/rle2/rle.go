package rle2

func Compress(input []byte) []byte {
	if len(input) == 0 {
		return input
	}

	output := make([]byte, 0, len(input))

	cursor := input[0]
	count := byte(1)
	isSame := true
	for i := 1; i < len(input); i++ {
		if isSame {
			if cursor == input[i] && count < 127 {
				count++
			} else if count > 1 {
				output = append(output, count, cursor)
				count = 1
			} else {
				count++
				isSame = false
			}
		} else {
			if cursor != input[i] && (i == len(input)-1 || i < len(input)-1 && input[i] != input[i+1]) && count < 128 {
				count++
			} else if count > 1 {
				output = append(output, byte(-int8(count)))
				for j := i - int(count); j < i; j++ {
					output = append(output, input[j])
				}
				count = 1
			} else {
				count++
				isSame = true
			}
		}

		cursor = input[i]
	}

	if isSame {
		output = append(output, count, cursor)
	} else {
		output = append(output, byte(-int8(count)))
		for j := len(input) - int(count); j < len(input); j++ {
			output = append(output, input[j])
		}
	}

	return output
}

func Decompress(input []byte) []byte {
	if len(input) == 0 {
		return input
	}
	if len(input) < 2 {
		panic("unexpected length of input")
	}

	output := make([]byte, 0, len(input))
	i := 0
	for i < len(input)-1 {
		count := input[i]
		cursor := input[i+1]
		if count <= 127 {
			for j := 0; j < int(count); j++ {
				output = append(output, cursor)
			}
			i += 2
		} else {
			i++
			for j := 0; j < int(byte(-int8(count))); j++ {
				output = append(output, input[i])
				i++
			}
		}
	}

	return output
}

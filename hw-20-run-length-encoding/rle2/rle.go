package rle2

func CompressBytes(in []byte) []byte {
	if len(in) == 0 {
		return in
	}
	out := make([]byte, 0, len(in))

	for i := 0; i < len(in); {
		sameCount := 1
		for j := i + 1; j < len(in) && in[i] == in[j] && sameCount < 127; j++ {
			sameCount++
		}
		if sameCount > 1 {
			out = append(out, byte(sameCount), in[i])
			i += sameCount

			continue
		}

		diffCount := 1
		for j := i; j < len(in)-1 && in[j] != in[j+1] && diffCount < 128; j++ {
			if j < len(in)-2 && in[j+1] == in[j+2] {
				break
			}
			diffCount++
		}
		if diffCount == 1 {
			out = append(out, 1, in[i])
		} else {
			out = append(out, byte(-int8(diffCount)))
			for j := i; j < i+diffCount; j++ {
				out = append(out, in[j])
			}
		}

		i += diffCount
	}

	return out
}

func DecompressBytes(in []byte) []byte {
	if len(in) == 0 {
		return in
	}
	if len(in) < 2 {
		panic("unexpected length of input")
	}

	out := make([]byte, 0, len(in))
	for i := 0; i < len(in)-1; {
		count := in[i]
		cursor := in[i+1]
		if count <= 127 {
			for j := 0; j < int(count); j++ {
				out = append(out, cursor)
			}
			i += 2
		} else {
			i++
			for j := 0; j < int(byte(-int8(count))); j++ {
				out = append(out, in[i])
				i++
			}
		}
	}

	return out
}

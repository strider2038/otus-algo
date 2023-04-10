package rle

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

const chunkSize = 4096

func Compress(input io.Reader, output io.Writer) error {
	chunk := make([]byte, chunkSize)

	do := true
	for do {
		n, err := input.Read(chunk)
		if errors.Is(err, io.EOF) {
			do = false
		} else if err != nil {
			return fmt.Errorf("read chunk: %w", err)
		}

		compressed := CompressBytes(chunk[0:n])
		if _, err := output.Write(compressed); err != nil {
			return fmt.Errorf("write chunk: %w", err)
		}
	}

	return nil
}

func Decompress(input io.Reader, output io.Writer) error {
	in := bufio.NewReader(input)

	for {
		count, err := in.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return fmt.Errorf("read count: %w", err)
		}

		cursor, err := in.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return fmt.Errorf("unexpected EOF")
			}
			return fmt.Errorf("read byte: %w", err)
		}

		for i := 0; i < int(count); i++ {
			if _, err := output.Write([]byte{cursor}); err != nil {
				return fmt.Errorf("write byte: %w", err)
			}
		}
	}
}

package rle

import (
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
	chunk := make([]byte, chunkSize)

	do := true
	for do {
		n, err := input.Read(chunk)
		if errors.Is(err, io.EOF) {
			do = false
		} else if err != nil {
			return fmt.Errorf("read chunk: %w", err)
		}

		decompressed := DecompressBytes(chunk[0:n])
		if _, err := output.Write(decompressed); err != nil {
			return fmt.Errorf("write chunk: %w", err)
		}
	}

	return nil
}

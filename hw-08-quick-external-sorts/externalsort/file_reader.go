package externalsort

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type IntReader struct {
	file    *os.File
	scanner *bufio.Scanner
}

func ReadIntegersFromFile(filename string) (*IntReader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	return &IntReader{
		file:    file,
		scanner: bufio.NewScanner(file),
	}, nil
}

func (reader *IntReader) Next() (int, error) {
	if !reader.scanner.Scan() {
		return 0, ErrEndOfList
	}

	number, err := strconv.Atoi(reader.scanner.Text())
	if err != nil {
		return 0, fmt.Errorf("parse int: %w", err)
	}

	return number, nil
}

func (reader *IntReader) Close() error {
	if reader == nil || reader.file == nil {
		return nil
	}

	return reader.file.Close()
}

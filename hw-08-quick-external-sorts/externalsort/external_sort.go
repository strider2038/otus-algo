package externalsort

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/strider2038/otus-algo/hw-08-quick-external-sorts/sort"
)

func SortFileV1(filename string, chunkSize int) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	chunkDir := filepath.Dir(filename)

	scanner := bufio.NewScanner(file)
	chunkCount := 0
	for {
		chunk, size, err := readChunk(chunkSize, scanner)
		if err != nil {
			return err
		}
		if size == 0 {
			break
		}

		chunk = sort.QuickMiddle(chunk)
		chunkName := fmt.Sprintf("%s/chunk_%d.txt", chunkDir, chunkCount)
		if err := writeChunk(chunk, size, chunkName); err != nil {
			return fmt.Errorf("write chunk %d: %w", chunkCount, err)
		}

		chunkCount++
	}

	files := make([]*os.File, chunkCount)
	defer func() {
		for i := 0; i < len(files); i++ {
			files[i].Close()
		}
	}()
	scanners := make([]*bufio.Scanner, chunkCount)
	for i := 0; i < chunkCount; i++ {
		chunkName := fmt.Sprintf("%s/chunk_%d.txt", chunkDir, chunkCount)
		file, err := os.Open(chunkName)
		if err != nil {
			return fmt.Errorf("open chunk: %w", err)
		}
		scanners[i] = bufio.NewScanner(file)
	}

	chunkNumbers := make([]int, chunkCount)
	eofs := make([]bool, chunkCount)
	for i := 0; i < chunkCount; i++ {
		eofs[i] = scanners[i].Scan()
		if !eofs[i] {
			chunkNumbers[i], err = strconv.Atoi(scanners[i].Text())
			if err != nil {
				return fmt.Errorf("read number: %w", err)
			}
		}
	}

	for {
		stop := true
		min := 0
		for i, number := range chunkNumbers {
			stop = stop && eofs[i]
			if !eofs[i] {
				min = number
			}
		}
		if stop {
			break
		}

		for i, number := range chunkNumbers {
			if number < min {
				min = number
			}
		}

	}

	return nil
}

func readChunk(chunkSize int, scanner *bufio.Scanner) ([]int, int, error) {
	chunk := make([]int, chunkSize)
	for i := 0; i < chunkSize; i++ {
		if !scanner.Scan() {
			return chunk, i, nil
		}
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, 0, fmt.Errorf("parse int: %w", err)
		}
		chunk[i] = number
	}

	return chunk, chunkSize, nil
}

func writeChunk(numbers []int, size int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	for i := 0; i < size; i++ {
		s := strconv.Itoa(numbers[i])
		if _, err := io.WriteString(file, s+"\n"); err != nil {
			return fmt.Errorf(`write to file "%s": %w`, filename, err)
		}
	}

	return nil
}

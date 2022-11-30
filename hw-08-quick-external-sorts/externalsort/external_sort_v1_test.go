package externalsort_test

import (
	"fmt"
	"io"
	"io/fs"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/strider2038/otus-algo/hw-08-quick-external-sorts/externalsort"
)

const TestFile = "testdata/items.txt"

func TestSort(t *testing.T) {
	tests := []struct {
		name string
		sort func(filename string, chunkSize int) error
	}{
		{name: "variant 1", sort: externalsort.SortFileV1},
	}
	paramsList := []struct {
		numbersCount int
		maxNumber    int
	}{
		{numbersCount: 100, maxNumber: 10},
		{numbersCount: 1_000, maxNumber: 10},
		{numbersCount: 10_000, maxNumber: 10},
		{numbersCount: 100_000, maxNumber: 10},
		{numbersCount: 1_000_000, maxNumber: 10},
		{numbersCount: 100, maxNumber: 100},
		{numbersCount: 1_000, maxNumber: 1_000},
		{numbersCount: 10_000, maxNumber: 10_000},
		{numbersCount: 100_000, maxNumber: 100_000},
		{numbersCount: 1_000_000, maxNumber: 1_000_000},
	}
	for _, test := range tests {
		for _, params := range paramsList {
			testName := fmt.Sprintf("%s, n=%d, t=%d", test.name, params.numbersCount, params.maxNumber)
			t.Run(testName, func(t *testing.T) {
				if err := os.RemoveAll("testdata"); err != nil {
					t.Fatal("remove testdata:", err)
				}
				if err := os.MkdirAll("testdata", fs.ModePerm); err != nil {
					t.Fatal("make testdata:", err)
				}
				if err := GenerateRandomDataFile(params.numbersCount, params.maxNumber, TestFile); err != nil {
					t.Fatal("generate file:", err)
				}

				if err := test.sort(TestFile, params.numbersCount/10); err != nil {
					t.Fatal("sort:", err)
				}

				items, err := ReadNumbers(TestFile)
				if err != nil {
					t.Fatal("read file:", err)
				}
				isSorted := sort.SliceIsSorted(items, func(i, j int) bool {
					return items[i] < items[j]
				})
				if !isSorted {
					t.Error("file is not sorted")
				}
			})
		}
	}
}

func GenerateRandomDataFile(n, t int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		number := strconv.Itoa(rand.Intn(t))
		if _, err := io.WriteString(file, number+"\n"); err != nil {
			return fmt.Errorf(`write to file "%s": %w`, filename, err)
		}
	}

	return nil
}

func ReadNumbers(filename string) ([]int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	numbers := make([]int, 0, len(lines))

	for i, line := range lines {
		if line == "" {
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("parse line %d: %w", i, err)
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

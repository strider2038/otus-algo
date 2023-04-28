package bloom_test

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"bloom/bloom"
	bitsbloom "github.com/bits-and-blooms/bloom/v3"
	"github.com/strider2038/otus-algo/datatesting"
)

func TestFilter_Contains(t *testing.T) {
	presentValues := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
	}
	absentValues := []string{
		"foo",
		"bar",
		"baz",
	}

	filter := bloom.NewFilter(10, 0.01)
	for _, value := range presentValues {
		filter.Add(value)
	}

	for _, value := range presentValues {
		datatesting.AssertTrue(t, filter.Contains(value))
	}
	for _, value := range absentValues {
		datatesting.AssertFalse(t, filter.Contains(value))
	}
}

func TestBigDataSet(t *testing.T) {
	const rowsCount = 500_000
	const maxValues = 100_000
	presentValues := make([]string, 0, maxValues)
	absentValues := make([]string, 0, maxValues)
	readDataSet(t, func(value string, isPresent bool) {
		if isPresent {
			if len(presentValues) < maxValues {
				presentValues = append(presentValues, value)
			}
		} else if len(absentValues) < maxValues {
			absentValues = append(absentValues, value)
		}
	})

	possibilities := []float64{0.1, 0.01, 0.001}
	for _, p := range possibilities {
		t.Run(fmt.Sprintf("%f", p), func(t *testing.T) {
			filter := bloom.NewFilter(rowsCount, p)
			count := 0
			readDataSet(t, func(value string, isPresent bool) {
				if isPresent {
					filter.Add(value)
					count++
				}
			})

			totalPositives := 0
			falsePositives := 0
			for _, value := range presentValues {
				contains := filter.Contains(value)
				datatesting.AssertTrue(t, contains)
				if contains {
					totalPositives++
				}
			}
			for _, value := range absentValues {
				if filter.Contains(value) {
					totalPositives++
					falsePositives++
				}
			}
			t.Log("total count:", count)
			t.Log("total positives:", totalPositives)
			t.Log("false positives:", falsePositives)
			t.Log("false positives rate:", float64(falsePositives)/float64(totalPositives))
			t.Log("size (KB):", filter.SizeBytes()/1024)
		})
	}
}

func TestBigDataSet_ForeignAlgo(t *testing.T) {
	const rowsCount = 500_000
	const maxValues = 100_000
	presentValues := make([]string, 0, maxValues)
	absentValues := make([]string, 0, maxValues)
	readDataSet(t, func(value string, isPresent bool) {
		if isPresent {
			if len(presentValues) < maxValues {
				presentValues = append(presentValues, value)
			}
		} else if len(absentValues) < maxValues {
			absentValues = append(absentValues, value)
		}
	})

	possibilities := []float64{0.1, 0.01, 0.001}
	for _, p := range possibilities {
		t.Run(fmt.Sprintf("%f", p), func(t *testing.T) {
			filter := bitsbloom.NewWithEstimates(rowsCount, p)
			count := 0
			readDataSet(t, func(value string, isPresent bool) {
				if isPresent {
					filter.AddString(value)
					count++
				}
			})

			totalPositives := 0
			falsePositives := 0
			for _, value := range presentValues {
				contains := filter.TestString(value)
				datatesting.AssertTrue(t, contains)
				if contains {
					totalPositives++
				}
			}
			for _, value := range absentValues {
				if filter.TestString(value) {
					totalPositives++
					falsePositives++
				}
			}
			t.Log("total count:", count)
			t.Log("total positives:", totalPositives)
			t.Log("false positives:", falsePositives)
			t.Log("false positives rate:", float64(falsePositives)/float64(totalPositives))
			t.Log("size (KB):", 16+len(filter.BitSet().Bytes())*8/1024)
		})
	}
}

func readDataSet(t *testing.T, f func(value string, isPresent bool)) {
	file, err := os.Open("./../testdata/urldata.csv")
	if os.IsNotExist(err) {
		t.Skip("dataset not found, skipping test")
	}
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	rows := csv.NewReader(file)
	rows.Read()

	for {
		row, err := rows.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			continue
		}
		if len(row) < 2 {
			t.Fatal("unexpected number of records")
		}
		f(row[0], row[1] == "good")
	}
}

package bloom_test

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"bloom/bloom"
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
	file, err := os.Open("./../testdata/urldata.csv")
	if os.IsNotExist(err) {
		t.Skip("dataset not found, skipping test")
	}
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	data := csv.NewReader(file)
	data.Read()

	const maxValues = 10000
	presentValues := make([]string, 0, maxValues)
	absentValues := make([]string, 0, maxValues)

	count := 0
	goodCount := 0
	badCount := 0
	filter := bloom.NewFilter(500_000, 0.01)
	for {
		row, err := data.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println(count, err)
			continue
		}
		if len(row) < 2 {
			t.Fatal("unexpected number of records")
		}
		if row[1] == "good" {
			filter.Add(row[0])
			count++
			if len(presentValues) < maxValues {
				presentValues = append(presentValues, row[0])
			}
			goodCount++
		} else {
			badCount++
			if len(absentValues) < maxValues {
				absentValues = append(absentValues, row[0])
			}
		}
	}

	fmt.Println(goodCount, badCount)

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
	t.Log("false positives rate:", float64(falsePositives)/float64(totalPositives)*100)
}

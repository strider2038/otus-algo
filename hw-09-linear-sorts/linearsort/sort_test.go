package linearsort_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-08-quick-external-sorts/sort"
	"github.com/strider2038/otus-algo/hw-09-linear-sorts/linearsort"
)

type Solver func(items []int) []int

func (s Solver) Solve(t *testing.T, input, output []string) {
	if len(input) < 2 || len(output) < 1 {
		t.Fatal(datatesting.ErrNotEnoughArguments)
	}

	unsorted, err := datatesting.ParseIntArray(input[1])
	if err != nil {
		t.Fatalf("parse input array: %v", err)
	}
	wantSorted, err := datatesting.ParseIntArray(output[0])
	if err != nil {
		t.Fatalf("parse output array: %v", err)
	}

	start := time.Now()
	gotSorted := s(unsorted)
	t.Log("elapsed time:", time.Since(start).String())

	datatesting.AssertEqualArrays(t, wantSorted, gotSorted)
}

func TestSortTable(t *testing.T) {
	tests := []struct {
		name         string
		sort         func(items []int) []int
		randomLimit  int
		digitsLimit  int
		sortedLimit  int
		reverseLimit int
	}{
		{
			name: "quick right",
			sort: sort.QuickRight[int],
		},
	}
	arrayTypes := []struct {
		name string
		dir  string
	}{
		{name: "random", dir: "0.random"},
		{name: "digits", dir: "1.digits"},
		{name: "sorted", dir: "2.sorted"},
		{name: "reverse", dir: "3.revers"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, arrayType := range arrayTypes {
				t.Run(arrayType.name, func(t *testing.T) {
					limit := 0
					switch arrayType.name {
					case "random":
						limit = test.randomLimit
					case "digits":
						limit = test.digitsLimit
					case "sorted":
						limit = test.sortedLimit
					case "reverse":
						limit = test.reverseLimit
					}
					runner := datatesting.NewRunner(
						datatesting.WithWorkdir(fmt.Sprintf("./../../testdata/sortdata/%s/", arrayType.dir)),
						datatesting.WithLimit(limit),
					)
					runner.Run(t, Solver(test.sort))
				})
			}
		})
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name string
		sort func(items []int) []int
	}{
		{
			name: "bucket",
			sort: linearsort.BucketSort,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			items := []int{325, 112, 234, 355, 404, 376, 702, 801, 997, 101}
			wantItems := []int{101, 112, 234, 325, 355, 376, 404, 702, 801, 997}

			got := test.sort(items)

			datatesting.AssertEqualArrays(t, wantItems, got)
		})
	}
}

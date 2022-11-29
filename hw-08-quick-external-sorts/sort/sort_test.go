package sort_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-08-quick-external-sorts/sort"
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
			name:         "quick right",
			sort:         sort.QuickRight[int],
			reverseLimit: 6,
		},
		{
			name: "quick middle",
			sort: sort.QuickMiddle[int],
		},
		{
			name:         "quick lomuto",
			sort:         sort.QuickLomuto[int],
			digitsLimit:  6,
			reverseLimit: 6,
		},
		{
			name: "merge",
			sort: sort.Merge[int],
		},
		{
			name: "merge with buffer",
			sort: sort.MergeWithBuffer[int],
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
			name: "quick right",
			sort: sort.QuickRight[int],
		},
		{
			name: "quick middle",
			sort: sort.QuickMiddle[int],
		},
		{
			name: "quick lomuto",
			sort: sort.QuickLomuto[int],
		},
		{
			name: "merge",
			sort: sort.Merge[int],
		},
		{
			name: "merge with buffer",
			sort: sort.MergeWithBuffer[int],
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			items := []int{3, 1, 2, 5, 4, 9, 7, 8, 10, 6}
			wantItems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

			got := test.sort(items)

			datatesting.AssertEqualArrays(t, wantItems, got)
		})
	}
}

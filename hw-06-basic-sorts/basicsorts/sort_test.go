package basicsorts_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-06-basic-sorts/basicsorts"
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
		name string
		sort func(items []int) []int
	}{
		{
			name: "bubble",
			sort: basicsorts.Bubble[int],
		},
		{
			name: "insertion",
			sort: basicsorts.Insertion[int],
		},
		{
			name: "shell",
			sort: basicsorts.Shell[int],
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
					runner := datatesting.NewRunner(
						datatesting.WithWorkdir(fmt.Sprintf("./../../testdata/sortdata/%s/", arrayType.dir)),
						datatesting.WithLimit(5),
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
			name: "bubble",
			sort: basicsorts.Bubble[int],
		},
		{
			name: "insertion",
			sort: basicsorts.Insertion[int],
		},
		{
			name: "shell",
			sort: basicsorts.Shell[int],
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			items := []int{3, 1, 2, 5, 4, 9, 7, 8, 10, 6}
			wantItems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

			got := test.sort(items)

			if len(wantItems) != len(got) {
				t.Fatalf("different length: want %d, got %d", len(wantItems), len(got))
			}
			for i := 0; i < len(wantItems); i++ {
				if wantItems[i] != got[i] {
					t.Errorf("different items at %d: want %d, got %d", i, wantItems[i], got[i])
				}
			}
		})
	}
}

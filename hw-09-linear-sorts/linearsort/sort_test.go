package linearsort_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-09-linear-sorts/linearsort"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name      string
		sort      func(items []int) []int
		items     []int
		wantItems []int
	}{
		{
			name:      "bucket",
			sort:      linearsort.BucketSort[int],
			items:     []int{14472, 51985, 7418, 25863, 54675, 25593, 43968, 60387, 59742, 25504},
			wantItems: []int{7418, 14472, 25504, 25593, 25863, 43968, 51985, 54675, 59742, 60387},
		},
		{
			name:      "bucket",
			sort:      linearsort.BucketSort[int],
			items:     []int{325, 112, 234, 355, 350, 404, 376, 702, 801, 997, 101},
			wantItems: []int{101, 112, 234, 325, 350, 355, 376, 404, 702, 801, 997},
		},
		{
			name:      "counting",
			sort:      func(items []int) []int { return linearsort.CountingSort(items, linearsort.Max(items)) },
			items:     []int{3, 1, 2, 5, 2, 7, 4, 9, 5, 7, 8, 10, 6},
			wantItems: []int{1, 2, 2, 3, 4, 5, 5, 6, 7, 7, 8, 9, 10},
		},
		{
			name:      "radix",
			sort:      func(items []int) []int { return linearsort.RadixSort(items, linearsort.Max(items), 3) },
			items:     []int{325, 112, 234, 355, 404, 376, 702, 801, 997, 101},
			wantItems: []int{101, 112, 234, 325, 355, 376, 404, 702, 801, 997},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.sort(test.items)

			datatesting.AssertEqualArrays(t, test.wantItems, got)
		})
	}
}

func TestSortRandom(t *testing.T) {
	// Заменить на 9 для теста на миллиарде чисел.
	const maxPower = 7

	tests := []struct {
		name string
		sort func(items []uint16) []uint16
	}{
		{
			name: "bucket",
			sort: linearsort.BucketSort[uint16],
		},
		{
			name: "counting",
			sort: func(items []uint16) []uint16 {
				return linearsort.CountingSort(items, 0xFFFF)
			},
		},
		{
			name: "radix (4 bit)",
			sort: func(items []uint16) []uint16 {
				return linearsort.RadixSort(items, 0xFFFF, 4)
			},
		},
		{
			name: "radix (8 bit)",
			sort: func(items []uint16) []uint16 {
				return linearsort.RadixSort(items, 0xFFFF, 8)
			},
		},
		{
			name: "radix (16 bit)",
			sort: func(items []uint16) []uint16 {
				return linearsort.RadixSort(items, 0xFFFF, 16)
			},
		},
	}
	for _, test := range tests {
		count := 10
		for i := 2; i <= maxPower; i++ {
			count *= 10
			t.Run(fmt.Sprintf("%s, n=%d", test.name, count), func(t *testing.T) {
				numbers := GenerateRandomNumbers[uint16](count, 0xFFFF)

				start := time.Now()
				sorted := test.sort(numbers)
				t.Log("elapsed time:", time.Since(start).String())

				isSorted := sort.SliceIsSorted(sorted, func(i, j int) bool {
					return sorted[i] < sorted[j]
				})
				if !isSorted {
					t.Error("numbers are not sorted")
				}
			})
		}
	}
}

func GenerateRandomNumbers[T linearsort.Sortable](n int, max T) []T {
	rand.Seed(time.Now().UnixNano())

	numbers := make([]T, n)
	for i := 0; i < n; i++ {
		numbers[i] = T(rand.Intn(int(max)))
	}

	return numbers
}

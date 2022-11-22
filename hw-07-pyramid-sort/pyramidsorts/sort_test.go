package pyramidsorts_test

import (
	"testing"

	"github.com/strider2038/otus-algo/hw-07-pyramid-sort/pyramidsorts"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name string
		sort func(items []int) []int
	}{
		{
			name: "selection",
			sort: pyramidsorts.Selection[int],
		},
		{
			name: "heap",
			sort: pyramidsorts.Heap[int],
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

package queue_test

import (
	"testing"

	"github.com/strider2038/otus-algo/hw-04-basic-data-structures/queue"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name      string
		items     []int
		skipN     int
		wantItem  int
		wantSize  int
		wantError error
	}{
		{
			name:      "empty queue",
			wantError: queue.ErrEmpty,
		},
		{
			name:     "single queue",
			items:    []int{1},
			wantItem: 1,
		},
		{
			name:     "filled queue, first element",
			items:    []int{1, 2, 3, 4},
			wantItem: 1,
			wantSize: 3,
		},
		{
			name:     "filled queue, mid element",
			items:    []int{1, 2, 3},
			skipN:    1,
			wantItem: 2,
			wantSize: 1,
		},
		{
			name:     "filled queue, last element",
			items:    []int{1, 2, 3, 4},
			skipN:    3,
			wantItem: 4,
			wantSize: 0,
		},
		{
			name:      "filled queue, all elements skipped",
			items:     []int{1, 2, 3, 4},
			skipN:     4,
			wantError: queue.ErrEmpty,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := queue.Queue[int]{}
			for _, item := range test.items {
				q.Enqueue(item)
			}
			for i := 0; i < test.skipN; i++ {
				q.Dequeue()
			}

			got, err := q.Dequeue()

			if test.wantError != nil {
				if test.wantError != err {
					t.Fatalf("want error %s, got %s", test.wantError, err)
				}
				return
			}
			if err != nil {
				t.Fatal("unexpected error:", err)
			}
			if test.wantSize != q.Size() {
				t.Errorf("unexpected size: want %d, got %d", test.wantSize, q.Size())
			}
			if test.wantItem != got {
				t.Errorf("unexpected last item: want %d, got %d", test.wantItem, got)
			}
		})
	}
}

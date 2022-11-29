package queue_test

import (
	"testing"

	"github.com/strider2038/otus-algo/hw-04-basic-data-structures/queue"
)

func TestPriorityQueue(t *testing.T) {
	tests := []struct {
		name      string
		gotQueue  func() *queue.PriorityQueue[int]
		wantItem  int
		wantSize  int
		wantError error
	}{
		{
			name: "empty queue",
			gotQueue: func() *queue.PriorityQueue[int] {
				return &queue.PriorityQueue[int]{}
			},
			wantError: queue.ErrEmpty,
		},
		{
			name: "single element queue",
			gotQueue: func() *queue.PriorityQueue[int] {
				q := &queue.PriorityQueue[int]{}
				q.Enqueue(1, 1)
				return q
			},
			wantItem: 1,
		},
		{
			name: "single priority queue",
			gotQueue: func() *queue.PriorityQueue[int] {
				q := &queue.PriorityQueue[int]{}
				q.Enqueue(1, 1)
				q.Enqueue(1, 2)
				return q
			},
			wantItem: 1,
			wantSize: 1,
		},
		{
			name: "multiple priority queue",
			gotQueue: func() *queue.PriorityQueue[int] {
				q := &queue.PriorityQueue[int]{}
				q.Enqueue(1, 1)
				q.Enqueue(1, 2)
				q.Enqueue(2, 3)
				q.Enqueue(2, 4)
				q.Enqueue(1, 5)
				return q
			},
			wantItem: 3,
			wantSize: 4,
		},
		{
			name: "insert into the middle",
			gotQueue: func() *queue.PriorityQueue[int] {
				q := &queue.PriorityQueue[int]{}
				q.Enqueue(1, 1)
				q.Enqueue(3, 2)
				q.Enqueue(5, 3)
				q.Enqueue(4, 4)
				q.Enqueue(2, 5)
				q.Dequeue()
				return q
			},
			wantItem: 4,
			wantSize: 3,
		},
		{
			name: "insert into the middle",
			gotQueue: func() *queue.PriorityQueue[int] {
				q := &queue.PriorityQueue[int]{}
				q.Enqueue(5, 1)
				q.Enqueue(3, 2)
				q.Enqueue(1, 3)
				q.Enqueue(4, 4)
				q.Enqueue(6, 5)
				q.Dequeue()
				q.Dequeue()
				return q
			},
			wantItem: 4,
			wantSize: 2,
		},
		{
			name: "emptying the queue",
			gotQueue: func() *queue.PriorityQueue[int] {
				q := &queue.PriorityQueue[int]{}
				q.Enqueue(5, 1)
				q.Enqueue(3, 2)
				q.Enqueue(1, 3)
				q.Dequeue()
				q.Dequeue()
				q.Dequeue()
				return q
			},
			wantError: queue.ErrEmpty,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := test.gotQueue()

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

package queue

import "errors"

var ErrEmpty = errors.New("queue is empty")

type Item[T any] struct {
	previous *Item[T]
	value    T
}

type Queue[T any] struct {
	front *Item[T]
	back  *Item[T]
	size  int
}

func NewQueue[T any](values ...T) *Queue[T] {
	q := &Queue[T]{}
	for _, value := range values {
		q.Enqueue(value)
	}
	return q
}

func (queue *Queue[T]) Size() int {
	return queue.size
}

// Enqueue - добавление в очередь. Класс сложности O(1).
func (queue *Queue[T]) Enqueue(value T) {
	item := &Item[T]{value: value}

	if queue.size == 0 {
		queue.front = item
	} else {
		queue.back.previous = item
	}

	queue.back = item
	queue.size++
}

// Dequeue - выборка из очереди. Класс сложности O(1).
func (queue *Queue[T]) Dequeue() (T, error) {
	if queue.size == 0 {
		var zero T

		return zero, ErrEmpty
	}

	value := queue.front.value
	queue.size--

	queue.front = queue.front.previous
	if queue.size == 0 {
		queue.back = nil
	}

	return value, nil
}

package graph

type Queue[T any] struct {
	front *QueueItem[T]
	back  *QueueItem[T]
	size  int
}

type QueueItem[T any] struct {
	previous *QueueItem[T]
	value    T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (queue *Queue[T]) Size() int {
	return queue.size
}

func (queue *Queue[T]) Enqueue(value T) {
	item := &QueueItem[T]{value: value}

	if queue.size == 0 {
		queue.front = item
	} else {
		queue.back.previous = item
	}

	queue.back = item
	queue.size++
}

func (queue *Queue[T]) Dequeue() (T, bool) {
	if queue.size == 0 {
		var zero T

		return zero, false
	}

	value := queue.front.value
	queue.size--

	queue.front = queue.front.previous
	if queue.size == 0 {
		queue.back = nil
	}

	return value, true
}

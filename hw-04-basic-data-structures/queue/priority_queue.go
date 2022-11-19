package queue

import "fmt"

type PriorityItem[T any] struct {
	priority int
	items    *Queue[T]
	next     *PriorityItem[T]
}

type PriorityQueue[T any] struct {
	head *PriorityItem[T]
	size int
}

func (queue *PriorityQueue[T]) Size() int {
	return queue.size
}

// Enqueue - добавление в очередь с приоритетом. Класс сложности O(N).
func (queue *PriorityQueue[T]) Enqueue(priority int, value T) {
	queue.size++

	// список пуст
	if queue.head == nil {
		queue.head = queue.newItem(priority, value)
		return
	}

	// ищем последний элемент в списке с подходящим приоритетом
	previous := queue.head
	item := queue.head
	for ; item.next != nil; item = item.next {
		// если заданный приоритет больше, то добавляем элемент перед следующим
		if priority > item.priority {
			newItem := queue.newItem(priority, value)
			newItem.next = item

			// новый элемент с максимальным приоритетом добавляем в начало списка
			if item == queue.head {
				queue.head = newItem
				return
			}

			previous.next = newItem
			return
		}

		// очередь с таким приоритетом существует
		if item.priority == priority {
			item.items.Enqueue(value)
			return
		}

		previous = item
	}

	// если дошли до конца, то добавляем элемент в конец списка
	item.next = queue.newItem(priority, value)
}

// Dequeue - выборка из очереди с максимальным приоритетом. Класс сложности O(1).
func (queue *PriorityQueue[T]) Dequeue() (T, error) {
	var zero T
	if queue.size == 0 {
		return zero, ErrEmpty
	}
	queue.size--

	for item := queue.head; item != nil; item = item.next {
		if item.items.Size() > 0 {
			if item.items.Size() == 1 {
				// оптимизация: убираем пустую очередь
				queue.head = item.next
			}

			return item.items.Dequeue()
		}
	}

	return zero, fmt.Errorf("unexpected error")
}

func (queue *PriorityQueue[T]) newItem(priority int, value T) *PriorityItem[T] {
	return &PriorityItem[T]{
		priority: priority,
		items:    NewQueue(value),
	}
}

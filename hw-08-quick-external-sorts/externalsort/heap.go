package externalsort

import "errors"

var ErrEndOfList = errors.New("end of list")

// MinHeap - куча с сортировкой от минимума. На вершине кучи - числа с минимальным размером.
type MinHeap struct {
	items []*HeapItem
	size  int
}

// GetMin - возвращает следующее минимальное число из k.
// Если все списки закончились, то возвращается ErrEndOfList.
// Так же может возвращаться ошибка извлечения следующего элемента из списка.
func (heap *MinHeap) GetMin() (int, error) {
	// элементов больше нет - сортировка закончилась
	if heap.size == 0 {
		return 0, ErrEndOfList
	}

	min := heap.items[0].value

	// выбираем следующий элемент из списка (файла)
	nextValue, err := heap.items[0].next.Next()
	// список (файл) закончился - убираем из кучи
	if errors.Is(err, ErrEndOfList) {
		// последний элемент ставим в начало
		heap.swap(0, heap.size-1)
		// уменьшаем размер кучи
		heap.size--
		// обновляем кучу
		heap.sort()

		return min, nil
	}
	// системная ошибка (например, при чтении файла)
	if err != nil {
		return 0, err
	}

	// подставляем значение
	heap.items[0].value = nextValue
	// обновляем кучу
	heap.sort()

	return min, nil
}

func (heap *MinHeap) sort() {
	for h := heap.size/2 - 1; h >= 0; h-- {
		heap.heapify(h)
	}
}

func (heap *MinHeap) heapify(root int) {
	parent := root
	left := 2*parent + 1
	right := left + 1
	if left < heap.size && heap.items[left].value < heap.items[parent].value {
		parent = left
	}
	if right < heap.size && heap.items[right].value < heap.items[parent].value {
		parent = right
	}
	if parent == root {
		return
	}
	heap.swap(parent, root)
	heap.heapify(parent)
}

func (heap *MinHeap) swap(i, j int) {
	heap.items[i], heap.items[j] = heap.items[j], heap.items[i]
}

// HeapItem - элемент кучи.
type HeapItem struct {
	// Текущее значение (прочитанное из файла)
	value int
	next  List
}

// List - абстракция списка для чтения следующего значения.
type List interface {
	Next() (int, error)
}

func NewMinHeapFromReaders(readers ...*IntReader) (*MinHeap, error) {
	heap := &MinHeap{items: make([]*HeapItem, len(readers))}

	for _, reader := range readers {
		value, err := reader.Next()
		// список (файл) пуст
		if errors.Is(err, ErrEndOfList) {
			continue
		}
		// ошибка чтения файла
		if err != nil {
			return nil, err
		}
		heap.items[heap.size] = &HeapItem{value: value, next: reader}
		heap.size++
	}

	heap.sort()

	return heap, nil
}

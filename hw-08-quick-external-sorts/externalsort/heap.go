package externalsort

import "errors"

var ErrEndOfList = errors.New("end of list")

type Heap struct {
	items []*HeapItem
	size  int
}

func (heap *Heap) GetMin() (int, error) {
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
		heap.normalize()

		return min, nil
	}
	// системная ошибка (например, при чтении файла)
	if err != nil {
		return 0, err
	}

	// подставляем значение
	heap.items[0].value = nextValue
	// обновляем кучу
	heap.normalize()

	return min, nil
}

func (heap *Heap) normalize() {
	for h := heap.size/2 - 1; h >= 0; h-- {
		heap.heapify(h)
	}
}

func (heap *Heap) heapify(root int) {
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

func (heap *Heap) swap(i, j int) {
	heap.items[i], heap.items[j] = heap.items[j], heap.items[i]
}

type HeapItem struct {
	// Текущее значение (прочитанное из файла)
	value int
	// Абстракция списка для чтения следующего значения
	next List
}

type List interface {
	Next() (int, error)
}

func NewHeapFromReaders(readers ...*IntReader) (*Heap, error) {
	heap := &Heap{items: make([]*HeapItem, len(readers))}

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

	heap.normalize()

	return heap, nil
}

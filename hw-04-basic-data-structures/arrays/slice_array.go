package arrays

type SliceArray[T any] struct {
	items []T
}

func NewSliceArray[T any](items ...T) *SliceArray[T] {
	return &SliceArray[T]{items: items}
}

func (arr *SliceArray[T]) Size() int {
	return len(arr.items)
}

func (arr *SliceArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(arr.items) {
		var zero T

		return zero, ErrIndexOutOfRange
	}

	return arr.items[index], nil
}

func (arr *SliceArray[T]) Set(index int, item T) error {
	if index < 0 || index >= len(arr.items) {
		return ErrIndexOutOfRange
	}

	arr.items[index] = item

	return nil
}

func (arr *SliceArray[T]) Add(item T) {
	arr.items = append(arr.items, item)
}

func (arr *SliceArray[T]) Insert(index int, item T) error {
	if index < 0 || index > len(arr.items) {
		return ErrIndexOutOfRange
	}
	if len(arr.items) == index {
		arr.items = append(arr.items, item)

		return nil
	}

	arr.items = append(arr.items[:index+1], arr.items[index:]...)
	arr.items[index] = item

	return nil
}

func (arr *SliceArray[T]) Remove(index int) (T, error) {
	if index < 0 || index >= len(arr.items) {
		var zero T

		return zero, ErrIndexOutOfRange
	}

	element := arr.items[index]

	arr.items = append(arr.items[:index], arr.items[index+1:]...)

	return element, nil
}

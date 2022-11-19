package arrays

type SingleArray[T any] struct {
	items []T
	size  int
}

func NewSingleArray[T any]() *SingleArray[T] {
	return &SingleArray[T]{}
}

func (arr *SingleArray[T]) Size() int {
	return arr.size
}

func (arr *SingleArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= arr.size {
		var zero T

		return zero, ErrIndexOutOfRange
	}

	return arr.items[index], nil
}

func (arr *SingleArray[T]) Set(index int, item T) error {
	if index < 0 || index >= arr.size {
		return ErrIndexOutOfRange
	}

	arr.items[index] = item

	return nil
}

func (arr *SingleArray[T]) Add(item T) {
	arr.increaseSize()
	arr.items[arr.size] = item
	arr.size++
}

func (arr *SingleArray[T]) increaseSize() {
	newArray := make([]T, len(arr.items)+1)
	for i := 0; i < len(arr.items); i++ {
		newArray[i] = arr.items[i]
	}
	arr.items = newArray
}

func (arr *SingleArray[T]) Insert(index int, item T) error {
	if index < 0 || index > arr.size {
		return ErrIndexOutOfRange
	}

	newArray := make([]T, arr.size+1)
	// копирование левой части
	for i := 0; i < index; i++ {
		newArray[i] = arr.items[i]
	}
	newArray[index] = item
	// копирование правой части
	for i := index + 1; i <= arr.size; i++ {
		newArray[i] = arr.items[i-1]
	}

	arr.items = newArray
	arr.size++

	return nil
}

func (arr *SingleArray[T]) Remove(index int) (T, error) {
	if index < 0 || index >= arr.size {
		var zero T

		return zero, ErrIndexOutOfRange
	}

	element := arr.items[index]

	// сдвиг второй половины массива влево
	for i := index + 1; i < arr.size; i++ {
		arr.items[i-1] = arr.items[i]
	}

	arr.size--

	return element, nil
}

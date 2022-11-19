package arrays

type VectorArray[T any] struct {
	vectorSize int
	items      []T
	size       int
}

func NewVectorArray[T any](vectorSize int) *VectorArray[T] {
	return &VectorArray[T]{vectorSize: vectorSize}
}

func (arr *VectorArray[T]) Size() int {
	return arr.size
}

func (arr *VectorArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= arr.size {
		var zero T

		return zero, ErrIndexOutOfRange
	}

	return arr.items[index], nil
}

func (arr *VectorArray[T]) Set(index int, item T) error {
	if index < 0 || index >= arr.size {
		return ErrIndexOutOfRange
	}

	arr.items[index] = item

	return nil
}

func (arr *VectorArray[T]) Add(item T) {
	if arr.size >= len(arr.items) {
		arr.increaseSize()
	}
	arr.items[arr.size] = item
	arr.size++
}

func (arr *VectorArray[T]) increaseSize() {
	newArray := arr.newResizedArray()
	for i := 0; i < arr.size; i++ {
		newArray[i] = arr.items[i]
	}
	arr.items = newArray
}

func (arr *VectorArray[T]) newResizedArray() []T {
	return make([]T, len(arr.items)+arr.vectorSize)
}

func (arr *VectorArray[T]) Insert(index int, item T) error {
	if index < 0 || index > arr.size {
		return ErrIndexOutOfRange
	}

	if arr.size < len(arr.items) {
		arr.shiftRight(index)
	} else {
		arr.shiftRightWithResize(index)
	}

	arr.items[index] = item
	arr.size++

	return nil
}

func (arr *VectorArray[T]) shiftRight(index int) {
	// сдвиг второй половины массива вправо
	for i := arr.size; i > index; i-- {
		arr.items[i] = arr.items[i-1]
	}
}

func (arr *VectorArray[T]) shiftRightWithResize(index int) {
	newArray := arr.newResizedArray()

	// копирование левой части
	for i := 0; i < index; i++ {
		newArray[i] = arr.items[i]
	}
	// копирование правой части
	for i := index + 1; i <= arr.size; i++ {
		newArray[i] = arr.items[i-1]
	}

	arr.items = newArray
}

func (arr *VectorArray[T]) Remove(index int) (T, error) {
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

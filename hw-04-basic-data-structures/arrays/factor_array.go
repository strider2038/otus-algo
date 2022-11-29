package arrays

type FactorArray[T any] struct {
	items  []T
	factor int
	size   int
}

func NewFactorArray[T any](factor, capacity int) *FactorArray[T] {
	if factor < 2 {
		panic("factor cannot be less than 2")
	}

	return &FactorArray[T]{factor: factor, items: make([]T, capacity)}
}

func (arr *FactorArray[T]) Size() int {
	return arr.size
}

func (arr *FactorArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= arr.size {
		var zero T

		return zero, ErrIndexOutOfRange
	}

	return arr.items[index], nil
}

func (arr *FactorArray[T]) Set(index int, item T) error {
	if index < 0 || index >= arr.size {
		return ErrIndexOutOfRange
	}

	arr.items[index] = item

	return nil
}

func (arr *FactorArray[T]) Add(item T) {
	if arr.size == len(arr.items) {
		arr.increaseSize()
	}
	arr.items[arr.size] = item
	arr.size++
}

func (arr *FactorArray[T]) increaseSize() {
	newArray := arr.newResizedArray()
	for i := 0; i < len(arr.items); i++ {
		newArray[i] = arr.items[i]
	}
	arr.items = newArray
}

func (arr *FactorArray[T]) newResizedArray() []T {
	newSize := 1
	if len(arr.items) > 0 {
		newSize = len(arr.items) * arr.factor
	}
	return make([]T, newSize)
}

func (arr *FactorArray[T]) Insert(index int, item T) error {
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

func (arr *FactorArray[T]) shiftRight(index int) {
	// сдвиг второй половины массива вправо
	for i := arr.size; i > index; i-- {
		arr.items[i] = arr.items[i-1]
	}
}

func (arr *FactorArray[T]) shiftRightWithResize(index int) {
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

func (arr *FactorArray[T]) Remove(index int) (T, error) {
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

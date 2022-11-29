package arrays

type MatrixArray[T any] struct {
	vectorSize int
	vectors    Array[Array[T]]
	size       int
}

func NewMatrixArray[T any](vectorSize int) *MatrixArray[T] {
	return &MatrixArray[T]{
		vectorSize: vectorSize,
		vectors:    NewSingleArray[Array[T]](),
	}
}

func (arr *MatrixArray[T]) Size() int {
	return arr.size
}

func (arr *MatrixArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= arr.size {
		var zero T
		return zero, ErrIndexOutOfRange
	}

	return arr.get(index)
}

func (arr *MatrixArray[T]) get(index int) (T, error) {
	items, err := arr.vectors.Get(index / arr.vectorSize)
	if err != nil {
		var zero T
		return zero, err
	}

	return items.Get(index % arr.vectorSize)
}

func (arr *MatrixArray[T]) Set(index int, item T) error {
	if index < 0 || index >= arr.size {
		return ErrIndexOutOfRange
	}

	return arr.set(index, item)
}

func (arr *MatrixArray[T]) set(index int, item T) error {
	items, err := arr.vectors.Get(index / arr.vectorSize)
	if err != nil {
		return err
	}

	return items.Set(index%arr.vectorSize, item)
}

func (arr *MatrixArray[T]) Add(item T) {
	arr.add(item)
	arr.size++
}

func (arr *MatrixArray[T]) add(item T) {
	if arr.size >= arr.vectors.Size()*arr.vectorSize {
		arr.vectors.Add(NewVectorArray[T](arr.vectorSize))
	}
	items, _ := arr.vectors.Get(arr.size / arr.vectorSize)
	items.Add(item)
}

func (arr *MatrixArray[T]) Insert(index int, item T) error {
	if arr.size >= arr.vectors.Size()*arr.vectorSize {
		arr.vectors.Add(NewVectorArray[T](arr.vectorSize))
	}
	// если указан последний элемент, то просто вызываем операцию добавления в конец
	if arr.size == index {
		arr.Add(item)

		return nil
	}

	// увеличиваем фактический размер массива добавляя в конец пустой элемент
	var zero T
	arr.add(zero)

	// сдвиг второй половины массива вправо
	for i := arr.size; i > index; i-- {
		element, err := arr.get(i - 1)
		if err != nil {
			return err
		}
		if err := arr.set(i, element); err != nil {
			return err
		}
	}

	if err := arr.set(index, item); err != nil {
		return err
	}

	arr.size++

	return nil
}

func (arr *MatrixArray[T]) Remove(index int) (T, error) {
	var zero T
	if index < 0 || index >= arr.size {
		return zero, ErrIndexOutOfRange
	}

	element, err := arr.get(index)
	if err != nil {
		return zero, ErrIndexOutOfRange
	}

	// сдвиг второй половины массива влево
	for i := index + 1; i < arr.size; i++ {
		item, err := arr.get(i)
		if err != nil {
			return zero, err
		}
		if err := arr.set(i-1, item); err != nil {
			return zero, err
		}
	}

	arr.size--

	return element, nil
}

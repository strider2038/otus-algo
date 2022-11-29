package arrays

type Array[T any] interface {
	Size() int
	Get(index int) (T, error)
	Set(index int, item T) error
	Add(item T)
	Insert(index int, item T) error
	Remove(index int) (T, error)
}

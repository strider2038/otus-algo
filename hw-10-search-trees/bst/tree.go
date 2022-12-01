package bst

type Key interface {
	~string |
		~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Node[K Key, V any] struct {
	Key   K
	Value V
	Left  *Node[K, V]
	Right *Node[K, V]
}

func (node *Node[K, V]) ForEach(f func(key K, value V) error) error {
	if node == nil {
		return nil
	}
	if err := node.Left.ForEach(f); err != nil {
		return err
	}
	if err := f(node.Key, node.Value); err != nil {
		return err
	}
	if err := node.Right.ForEach(f); err != nil {
		return err
	}

	return nil
}

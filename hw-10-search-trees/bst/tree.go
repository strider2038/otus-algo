package bst

import "errors"

var ErrNotFound = errors.New("not found")

type Value interface {
	~string |
		~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Tree[V Value] struct {
	root *Node[V]
}

func (tree *Tree[V]) Find(value V) (*Node[V], error) {
	node := tree.root.find(value)
	if node == nil {
		return nil, ErrNotFound
	}

	return node, nil
}

func (tree *Tree[V]) Insert(value V) {
	tree.root = tree.root.insert(value)
}

func (tree *Tree[V]) Remove(value V) {
	tree.root = tree.root.remove(value)
}

func (tree *Tree[V]) ForEach(f func(value V) error) error {
	return tree.root.forEach(f)
}

type Node[V Value] struct {
	Value V
	Left  *Node[V]
	Right *Node[V]
}

func (node *Node[V]) find(value V) *Node[V] {
	if node == nil {
		return nil
	}
	if value == node.Value {
		return node
	}
	if value < node.Value {
		return node.Left.find(value)
	}

	return node.Right.find(value)
}

func (node *Node[V]) insert(value V) *Node[V] {
	if node == nil {
		return &Node[V]{Value: value}
	}

	if value <= node.Value {
		node.Left = node.Left.insert(value)
	} else {
		node.Right = node.Right.insert(value)
	}

	return node
}

func (node *Node[V]) remove(value V) *Node[V] {
	if node == nil {
		return nil
	}

	// значение меньше текущего - идем влево
	if value < node.Value {
		node.Left = node.Left.remove(value)

		return node
	}
	// значение больше текущего - идем вправо
	if value > node.Value {
		node.Right = node.Right.remove(value)

		return node
	}

	// искомый узел найден

	// узел без детей - возвращаем nil (на вызывающей стороне ссылка будет стерта)
	if node.Left == nil && node.Right == nil {
		return nil
	}

	// левого узла нет - возвращаем правый узел
	if node.Left == nil {
		return node.Right
	}

	// правого узла нет - возвращаем левый узел
	if node.Right == nil {
		return node.Left
	}

	// ищем наименьший узел в правой части дерева
	smallestOnRight := node.Right
	for smallestOnRight != nil && smallestOnRight.Left != nil {
		smallestOnRight = smallestOnRight.Left
	}

	// заменяем текущий узел наименьшим из правой части
	node.Value = smallestOnRight.Value
	// удаляем наименьший узел из правой части
	node.Right = node.Right.remove(node.Value)

	return node
}

func (node *Node[V]) forEach(f func(value V) error) error {
	if node == nil {
		return nil
	}
	if err := node.Left.forEach(f); err != nil {
		return err
	}
	if err := f(node.Value); err != nil {
		return err
	}
	if err := node.Right.forEach(f); err != nil {
		return err
	}

	return nil
}

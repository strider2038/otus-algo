package avl

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
	Value  V
	height int
	left   *Node[V]
	right  *Node[V]
}

func (node *Node[V]) find(value V) *Node[V] {
	if node == nil {
		return nil
	}
	if value == node.Value {
		return node
	}
	if value < node.Value {
		return node.left.find(value)
	}

	return node.right.find(value)
}

func (node *Node[V]) insert(value V) *Node[V] {
	if node == nil {
		return &Node[V]{Value: value, height: 1}
	}

	if value <= node.Value {
		node.left = node.left.insert(value)
	} else {
		node.right = node.right.insert(value)
	}

	return node.rebalance()
}

func (node *Node[V]) remove(value V) *Node[V] {
	if node == nil {
		return nil
	}

	if value < node.Value { // значение меньше текущего - идем влево
		node.left = node.left.remove(value)
	} else if value > node.Value { // значение больше текущего - идем вправо
		node.right = node.right.remove(value)
		// искомый узел найден
	} else if node.left == nil && node.right == nil {
		// узел без детей - возвращаем nil (на вызывающей стороне ссылка будет стерта)
		node = nil
	} else if node.left == nil { // левого узла нет - возвращаем правый узел
		node = node.right
	} else if node.right == nil { // правого узла нет - возвращаем левый узел
		node = node.left
	} else {
		// ищем наименьший узел в правой части дерева
		smallestOnRight := node.right
		for smallestOnRight != nil && smallestOnRight.left != nil {
			smallestOnRight = smallestOnRight.left
		}

		// заменяем текущий узел наименьшим из правой части
		node.Value = smallestOnRight.Value
		// удаляем наименьший узел из правой части
		node.right = node.right.remove(node.Value)
	}

	return node.rebalance()
}

func (node *Node[V]) forEach(f func(value V) error) error {
	if node == nil {
		return nil
	}
	if err := node.left.forEach(f); err != nil {
		return err
	}
	if err := f(node.Value); err != nil {
		return err
	}
	if err := node.right.forEach(f); err != nil {
		return err
	}

	return nil
}

func (node *Node[V]) rebalance() *Node[V] {
	if node == nil {
		return nil
	}

	node.recalculateHeight()

	// проверка баланса дерева
	balance := node.left.getHeight() - node.right.getHeight()

	// вращаем влево, если высота левого поддерева больше правого
	if balance <= -2 {
		// если высота левого поддерева больше правого, то
		// выполняем большой поворот влево
		if node.right.left.getHeight() > node.right.right.getHeight() {
			return node.bigRotateLeft()
		}

		return node.smallRotateLeft()
	}

	if balance >= 2 {
		// если высота правого поддерева больше левого, то
		// выполняем большой поворот вправо
		if node.left.right.getHeight() > node.left.left.getHeight() {
			return node.bigRotateRight()
		}

		return node.smallRotateRight()
	}

	return node
}

func (node *Node[V]) bigRotateLeft() *Node[V] {
	node.right = node.right.smallRotateRight()

	return node.smallRotateLeft()
}

func (node *Node[V]) bigRotateRight() *Node[V] {
	node.left = node.left.smallRotateLeft()

	return node.smallRotateRight()
}

func (node *Node[V]) smallRotateLeft() *Node[V] {
	newRoot := node.right
	node.right = newRoot.left
	newRoot.left = node

	node.recalculateHeight()
	newRoot.recalculateHeight()

	return newRoot
}

func (node *Node[V]) smallRotateRight() *Node[V] {
	newRoot := node.left
	node.left = newRoot.right
	newRoot.right = node

	node.recalculateHeight()
	newRoot.recalculateHeight()

	return newRoot
}

func (node *Node[V]) recalculateHeight() {
	if node == nil {
		return
	}

	leftHeight := node.left.getHeight()
	rightHeight := node.right.getHeight()

	if leftHeight > rightHeight {
		node.height = leftHeight + 1
	} else {
		node.height = rightHeight + 1
	}
}

func (node *Node[V]) getHeight() int {
	if node == nil {
		return 0
	}

	return node.height
}

package obst

import (
	"errors"

	"github.com/strider2038/otus-algo/hw-08-quick-external-sorts/sort"
)

var ErrNotFound = errors.New("not found")

type Value interface {
	~string |
		~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type InputNode[V Value] struct {
	Value  V
	Weight float64
}

func NewV1[V Value](nodes ...InputNode[V]) *Tree[V] {
	tree := &Tree[V]{}

	nodes = sort.QuickSlice(nodes, func(i, j int) bool {
		return nodes[i].Weight < nodes[j].Weight
	})

	for _, node := range nodes {
		tree.insert(node.Value, node.Weight)
	}

	return tree
}

func NewV2[V Value](nodes ...InputNode[V]) *Tree[V] {
	tree := &Tree[V]{}

	nodes = sort.QuickSlice(nodes, func(i, j int) bool {
		return nodes[i].Value > nodes[j].Value
	})

	fillTree(tree, nodes, 0, len(nodes)-1)

	return tree
}

func fillTree[V Value](tree *Tree[V], nodes []InputNode[V], left, right int) {
	if left > right {
		return
	}

	// суммарный вес отрезка
	weight := 0.0
	for i := left; i <= right; i++ {
		weight += nodes[i].Weight
	}

	// поиск "центра тяжести"
	sum := 0.0
	i := left
	for ; i <= right; i++ {
		if sum < weight/2 && sum+nodes[i].Weight > weight/2 {
			break
		}
		sum += nodes[i].Weight
	}

	tree.insert(nodes[i].Value, nodes[i].Weight)
	fillTree(tree, nodes, left, i-1)
	fillTree(tree, nodes, i+1, right)
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

func (tree *Tree[V]) ForEach(f func(value V, weight float64) error) error {
	return tree.root.forEach(f)
}

func (tree *Tree[V]) insert(value V, weight float64) {
	tree.root = tree.root.insert(value, weight)
}

type Node[V Value] struct {
	Value  V
	Weight float64
	Left   *Node[V]
	Right  *Node[V]
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

func (node *Node[V]) insert(value V, weight float64) *Node[V] {
	if node == nil {
		return &Node[V]{Value: value, Weight: weight}
	}

	if value <= node.Value {
		node.Left = node.Left.insert(value, weight)
	} else {
		node.Right = node.Right.insert(value, weight)
	}

	return node
}

func (node *Node[V]) forEach(f func(value V, weight float64) error) error {
	if node == nil {
		return nil
	}
	if err := node.Left.forEach(f); err != nil {
		return err
	}
	if err := f(node.Value, node.Weight); err != nil {
		return err
	}
	if err := node.Right.forEach(f); err != nil {
		return err
	}

	return nil
}

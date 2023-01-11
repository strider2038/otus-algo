package triearray

const alphabetSize = 26

type Trie[V any] struct {
	root  Node[V]
	count int
}

type Node[V any] struct {
	children [alphabetSize]*Node[V]
	value    *V
}

func New[V any]() Trie[V] {
	return Trie[V]{}
}

func (t *Trie[V]) Count() int {
	return t.count
}

func (t *Trie[V]) Get(key string) V {
	v, _ := t.Find(key)

	return v
}

func (t *Trie[V]) Find(key string) (V, bool) {
	node := &t.root
	var zero V

	for _, char := range key {
		index := char - 'a'
		if node.children[index] == nil {
			return zero, false
		}
		node = node.children[index]
	}

	if node != nil && node.value != nil {
		return *node.value, true
	}

	return zero, false
}

func (t *Trie[V]) Put(key string, value V) {
	node := &t.root

	for _, char := range key {
		index := char - 'a'
		if node.children[index] == nil {
			node.children[index] = &Node[V]{}
		}
		node = node.children[index]
	}

	if node.value == nil {
		t.count++
	}

	node.value = &value
}

func (t *Trie[V]) Delete(key string) {
	node := &t.root

	for _, char := range key {
		index := char - 'a'
		if index < 0 || index >= alphabetSize {
			panic("index out of range")
		}
		if node.children[index] == nil {
			return
		}
		node = node.children[index]
	}

	node.value = nil
	t.count--
}

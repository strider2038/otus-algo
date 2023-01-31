package trie

const alphabetSize = 26

type Trie struct {
	root Node
}

type Node struct {
	// todo: use pointer
	children *[alphabetSize]*Node
	isEnd    bool
}

func New() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := &t.root

	for _, char := range word {
		index := char - 'a'
		if index < 0 || index >= alphabetSize {
			panic("index out of range")
		}
		if node.children == nil {
			node.children = &[alphabetSize]*Node{}
		}
		if node.children[index] == nil {
			node.children[index] = &Node{}
		}
		node = node.children[index]
	}

	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := &t.root

	for _, char := range word {
		index := char - 'a'
		if index < 0 || index >= alphabetSize {
			panic("index out of range")
		}
		if node.children == nil || node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}

	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := &t.root

	for _, char := range prefix {
		index := char - 'a'
		if index < 0 || index >= alphabetSize {
			panic("index out of range")
		}
		if node.children == nil || node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}

	return true
}

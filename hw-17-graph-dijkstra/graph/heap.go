package graph

// minHeap - куча с сортировкой от минимума. На вершине кучи - числа с минимальным размером.
type minHeap struct {
	items []*Direction
	size  int
}

func newMinHeap(capacity int) *minHeap {
	return &minHeap{items: make([]*Direction, capacity)}
}

func (heap *minHeap) Pop() (*Direction, bool) {
	if heap.size == 0 {
		return nil, false
	}

	min := heap.items[0]

	heap.items[0] = heap.items[heap.size-1]
	heap.size--
	heap.sort()

	return min, true
}

func (heap *minHeap) Insert(d *Direction) {
	heap.items[heap.size] = d
	heap.size++
	heap.sort()
}

func (heap *minHeap) sort() {
	for h := heap.size/2 - 1; h >= 0; h-- {
		heap.heapify(h)
	}
}

func (heap *minHeap) heapify(root int) {
	parent := root
	left := 2*parent + 1
	right := left + 1
	if left < heap.size && heap.items[left].Distance < heap.items[parent].Distance {
		parent = left
	}
	if right < heap.size && heap.items[right].Distance < heap.items[parent].Distance {
		parent = right
	}
	if parent == root {
		return
	}
	heap.swap(parent, root)
	heap.heapify(parent)
}

func (heap *minHeap) swap(i, j int) {
	heap.items[i], heap.items[j] = heap.items[j], heap.items[i]
}

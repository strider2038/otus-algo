package pyramidsorts

func Heap[T Sortable](items []T) []T {
	for h := len(items)/2 - 1; h >= 0; h-- {
		heapify(items, h, len(items))
	}
	for j := len(items) - 1; j > 0; j-- {
		swap(items, 0, j)
		heapify(items, 0, j)
	}

	return items
}

func heapify[T Sortable](items []T, root, size int) {
	parent := root
	left := 2*parent + 1
	right := left + 1
	if left < size && items[left] > items[parent] {
		parent = left
	}
	if right < size && items[right] > items[parent] {
		parent = right
	}
	if parent == root {
		return
	}
	swap(items, parent, root)
	heapify(items, parent, size)
}

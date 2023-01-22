package sort

func QuickSlice[T any](items []T, greater func(i, j int) bool) []T {
	quickSortSlice(items, 0, len(items)-1, greater)
	return items
}

func quickSortSlice[T any](items []T, left, right int, greater func(i, j int) bool) {
	if left >= right {
		return
	}
	middle := quickSortSliceSplit(items, left, right, greater)
	quickSortSlice(items, left, middle-1, greater)
	quickSortSlice(items, middle+1, right, greater)
}

func quickSortSliceSplit[T any](items []T, left, right int, greater func(i, j int) bool) int {
	pivot := right
	middle := left - 1

	for j := left; j <= right; j++ {
		if !greater(j, pivot) {
			middle++
			swap(items, middle, j)
		}
	}

	return middle
}

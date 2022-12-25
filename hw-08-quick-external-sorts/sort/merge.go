package sort

func Merge[T Sortable](items []T) []T {
	mergeSort(items, 0, len(items)-1)
	return items
}

func mergeSort[T Sortable](items []T, left, right int) {
	if left >= right {
		return
	}
	middle := (left + right) / 2
	mergeSort(items, left, middle)
	mergeSort(items, middle+1, right)
	merge(items, left, middle, right)
}

func merge[T Sortable](items []T, left, middle, right int) {
	temp := make([]T, right-left+1)
	a := left
	b := middle + 1
	i := 0
	for ; a <= middle && b <= right; i++ {
		if items[a] < items[b] {
			temp[i] = items[a]
			a++
		} else {
			temp[i] = items[b]
			b++
		}
	}
	for ; a <= middle; i++ {
		temp[i] = items[a]
		a++
	}
	for ; b <= right; i++ {
		temp[i] = items[b]
		b++
	}
	for j := left; j <= right; j++ {
		items[j] = temp[j-left]
	}
}

func MergeWithBuffer[T Sortable](items []T) []T {
	buffer := make([]T, len(items))
	mergeWithBufferSort(items, buffer, 0, len(items)-1)
	return items
}

func mergeWithBufferSort[T Sortable](items, buffer []T, left, right int) {
	if left >= right {
		return
	}
	middle := (left + right) / 2
	mergeWithBufferSort(items, buffer, left, middle)
	mergeWithBufferSort(items, buffer, middle+1, right)
	mergeWithBuffer(items, buffer, left, middle, right)
}

func mergeWithBuffer[T Sortable](items, buffer []T, left, middle, right int) {
	a := left
	b := middle + 1
	i := left
	for ; a <= middle && b <= right; i++ {
		if items[a] < items[b] {
			buffer[i] = items[a]
			a++
		} else {
			buffer[i] = items[b]
			b++
		}
	}
	for ; a <= middle; i++ {
		buffer[i] = items[a]
		a++
	}
	for ; b <= right; i++ {
		buffer[i] = items[b]
		b++
	}
	for j := left; j <= right; j++ {
		items[j] = buffer[j]
	}
}

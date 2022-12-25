package sort

// QuickRight - алгоритм быстрой сортировки, вариант разделения с выбором
// крайне правого значения как опорного.
func QuickRight[T Sortable](items []T) []T {
	quickRightSort(items, 0, len(items)-1)
	return items
}

func quickRightSort[T Sortable](items []T, left, right int) {
	if left >= right {
		return
	}
	middle := quickRightSplit(items, left, right)
	quickRightSort(items, left, middle)
	quickRightSort(items, middle+1, right)
}

func quickRightSplit[T Sortable](items []T, left, right int) int {
	pivot := items[right-1]
	i := left
	j := right
	for i <= j {
		for items[i] < pivot {
			i++
		}
		for items[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		swap(items, i, j)
		i++
		j--
	}

	return j
}

// QuickMiddle - алгоритм быстрой сортировки, вариант разделения с выбором
// середины отрезка. См. https://neerc.ifmo.ru/wiki/index.php?title=Быстрая_сортировка.
func QuickMiddle[T Sortable](items []T) []T {
	quickMiddleSort(items, 0, len(items)-1)
	return items
}

func quickMiddleSort[T Sortable](items []T, left, right int) {
	if left >= right {
		return
	}
	middle := quickMiddleSplit(items, left, right)
	quickMiddleSort(items, left, middle)
	quickMiddleSort(items, middle+1, right)
}

func quickMiddleSplit[T Sortable](items []T, left, right int) int {
	pivot := items[(left+right)/2]
	i := left
	j := right
	for i <= j {
		for items[i] < pivot {
			i++
		}
		for items[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		swap(items, i, j)
		i++
		j--
	}

	return j
}

// QuickLomuto - алгоритм быстрой сортировки, вариант разделения Lomuto.
// См. https://en.wikipedia.org/wiki/Quicksort.
func QuickLomuto[T Sortable](items []T) []T {
	quickLomutoSort(items, 0, len(items)-1)
	return items
}

func quickLomutoSort[T Sortable](items []T, left, right int) {
	if left >= right {
		return
	}
	middle := quickLomutoSplit(items, left, right)
	quickLomutoSort(items, left, middle-1)
	quickLomutoSort(items, middle+1, right)
}

func quickLomutoSplit[T Sortable](items []T, left, right int) int {
	pivot := items[right]
	middle := left - 1

	for j := left; j <= right; j++ {
		if items[j] <= pivot {
			middle++
			swap(items, middle, j)
		}
	}

	return middle
}

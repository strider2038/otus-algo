package basicsorts

func Insertion[T Sortable](items []T) []T {
	for i := 1; i < len(items); i++ {
		for j := i - 1; j >= 0 && items[j] > items[j+1]; j-- {
			swap(items, j, j+1)
		}
	}

	return items
}

func InsertionShift[T Sortable](items []T) []T {
	for i := 1; i < len(items); i++ {
		// запоминаем элемент для переноса
		k := items[i]
		// сдвигаем отсортированную часть массива, значения которой больше k
		j := 0
		for j = i - 1; j >= 0 && items[j] > k; j-- {
			items[j+1] = items[j]
		}
		// подставляем элемент в отсортированную часть
		items[j+1] = k
	}

	return items
}

func InsertionBinarySearch[T Sortable](items []T) []T {
	j := 0

	for i := 1; i < len(items); i++ {
		// запоминаем элемент для переноса
		k := items[i]
		// ищем индекс наименьшего элемента, который больше k
		// с помощью бинарного поиска
		p := binarySearch(items, k, 0, i-1)
		// сдвигаем отсортированную часть массива, значения которой больше k
		for j = i - 1; j >= p; j-- {
			items[j+1] = items[j]
		}
		// подставляем элемент в отсортированную часть
		items[j+1] = k
	}

	return items
}

func binarySearch[T Sortable](items []T, value T, low, high int) int {
	if high <= low {
		if value > items[low] {
			return low + 1
		}
		return low
	}

	mid := (low + high) / 2
	if value > items[mid] {
		return binarySearch(items, value, mid+1, high)
	}

	return binarySearch(items, value, low, mid-1)
}

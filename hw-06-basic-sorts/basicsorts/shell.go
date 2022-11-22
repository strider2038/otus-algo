package basicsorts

func Shell[T Sortable](items []T) []T {
	for gap := len(items) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(items); i++ {
			for j := i - gap; j >= 0 && items[j] > items[j+gap]; j -= gap {
				swap(items, j+gap, j)
			}
		}
	}

	return items
}

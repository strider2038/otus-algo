package basicsorts

func Insertion[T Sortable](items []T) []T {
	for i := 1; i < len(items); i++ {
		for j := i - 1; j >= 0 && items[j] > items[j+1]; j-- {
			swap(items, j, j+1)
		}
	}

	return items
}

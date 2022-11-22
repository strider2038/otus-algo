package basicsorts

func Bubble[T Sortable](items []T) []T {
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items)-1; j++ {
			if items[j] > items[j+1] {
				swap(items, j, j+1)
			}
		}
	}

	return items
}

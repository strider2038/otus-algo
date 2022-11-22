package pyramidsorts

func Selection[T Sortable](items []T) []T {
	for i := 0; i < len(items)-1; i++ {
		min := i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[min] {
				min = j
			}
		}

		t := items[i]
		items[i] = items[min]
		items[min] = t
	}

	return items
}

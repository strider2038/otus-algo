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

func BubbleOptimized[T Sortable](items []T) []T {
	for n := len(items); n > 0; {
		newN := 0
		for j := 0; j < n-1; j++ {
			if items[j] > items[j+1] {
				swap(items, j, j+1)
				newN = j + 1
			}
		}
		n = newN
	}

	return items
}

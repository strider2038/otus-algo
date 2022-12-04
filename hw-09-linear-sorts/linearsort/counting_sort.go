package linearsort

func CountingSort[T Sortable](items []T, max int) []T {
	counts := make([]int, max+1)
	for i := 0; i < len(items); i++ {
		counts[items[i]]++
	}

	for i := 1; i <= max; i++ {
		counts[i] += counts[i-1]
	}

	sorted := make([]T, len(items))
	for i := len(items) - 1; i >= 0; i-- {
		t := items[i]
		counts[t]--
		sorted[counts[t]] = t
	}

	return sorted
}

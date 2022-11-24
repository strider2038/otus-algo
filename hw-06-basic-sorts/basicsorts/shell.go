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

func ShellFrankLazarus[T Sortable](items []T) []T {
	for gap := len(items)/4*2 + 1; ; gap = gap/4*2 + 1 {
		for i := gap; i < len(items); i++ {
			for j := i - gap; j >= 0 && items[j] > items[j+gap]; j -= gap {
				swap(items, j+gap, j)
			}
		}
		if gap <= 1 {
			break
		}
	}

	return items
}

func ShellInsertion[T Sortable](items []T) []T {
	for gap := len(items) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(items); i++ {
			// сортировка вставками (оптимизация со сдвигом)
			k := items[i]

			j := 0
			for j = i; j >= gap && items[j-gap] > k; j -= gap {
				items[j] = items[j-gap]
			}

			items[j] = k
		}
	}

	return items
}

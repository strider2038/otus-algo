package linearsort

type Sortable interface {
	~string |
		~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func swap[T Sortable](items []T, i, j int) {
	t := items[i]
	items[i] = items[j]
	items[j] = t
}

func Max[T Sortable](items []T) T {
	var zero T
	if len(items) == 0 {
		return zero
	}

	max := items[0]

	for i := 1; i < len(items); i++ {
		if items[i] > max {
			max = items[i]
		}
	}

	return max
}

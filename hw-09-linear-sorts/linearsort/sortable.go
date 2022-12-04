package linearsort

type Sortable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
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

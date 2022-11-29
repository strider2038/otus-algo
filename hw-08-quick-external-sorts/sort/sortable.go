package sort

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

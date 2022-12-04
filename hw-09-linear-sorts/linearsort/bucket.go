package linearsort

func BucketSort[T Sortable](items []T) []T {
	n := len(items)
	max := Max(items)

	buckets := make([]Bucket[T], n)
	for i := 0; i < n; i++ {
		bucketIndex := n * int(items[i]) / (int(max) + 1)
		buckets[bucketIndex].Add(items[i])
	}

	i := 0
	for _, bucket := range buckets {
		bucket.ForEach(func(value T) {
			items[i] = value
			i++
		})
	}

	return items
}

type Bucket[T Sortable] struct {
	head *BucketItem[T]
}

type BucketItem[T Sortable] struct {
	next  *BucketItem[T]
	value T
}

func (bucket *Bucket[T]) Add(value T) {
	if bucket.head == nil {
		bucket.head = &BucketItem[T]{value: value}
		return
	}
	if value <= bucket.head.value {
		bucket.head = &BucketItem[T]{value: value, next: bucket.head}
		return
	}

	for node := bucket.head; ; node = node.next {
		if node.next == nil || value <= node.next.value {
			node.next = &BucketItem[T]{value: value, next: node.next}

			return
		}
	}
}

func (bucket *Bucket[T]) ForEach(f func(value T)) {
	for current := bucket.head; current != nil; current = current.next {
		f(current.value)
	}
}

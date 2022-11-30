package linearsort

func BucketSort(items []int) []int {
	n := len(items)
	max := Max(items)

	buckets := make([]Bucket, n)
	for i := 0; i < n; i++ {
		bucketIndex := n * items[i] / (max + 1)
		buckets[bucketIndex].Add(items[i])
	}

	return items
}

type Bucket struct {
	head *BucketItem
}

func (bucket *Bucket) Add(value int) {
	if bucket.head == nil {
		bucket.head = &BucketItem{value: value}
		return
	}

	prev := bucket.head
	current := bucket.head
	for ; current != nil; current = current.next {
		if value < current.value {
			if prev == bucket.head {
				bucket.head = &BucketItem{value: value, next: current}
			} else {
				prev.next = &BucketItem{value: value, next: current}
			}

			return
		}

		prev = current
	}

	if current == nil {
		prev.next = &BucketItem{value: value, next: current}
	}
}

type BucketItem struct {
	next  *BucketItem
	value int
}

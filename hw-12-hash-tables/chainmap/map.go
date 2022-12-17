package chainmap

import (
	"crypto/sha256"
	"encoding/binary"
)

const (
	initialSize   = 8
	maxLoadFactor = 0.85
)

type Map[V any] struct {
	count int
	items []*Item[V]
}

type Item[V any] struct {
	key   string
	value V
	next  *Item[V]
}

func (m *Map[V]) Get(key string) V {
	v, _ := m.Find(key)

	return v
}

func (m *Map[V]) Find(key string) (V, bool) {
	bucketNo := m.getBucketNo(key)

	for item := m.items[bucketNo]; item != nil; item = item.next {
		if item.key == key {
			return item.value, true
		}
	}

	var zero V

	return zero, false
}

func (m *Map[V]) Put(key string, value V) {
	m.rehash()

	if m.put(key, value) {
		m.count++
	}
}

func (m *Map[V]) Delete(key string) {
	bucketNo := m.getBucketNo(key)

	if m.items[bucketNo] == nil {
		return
	}

	if m.items[bucketNo].key == key {
		m.items[bucketNo] = m.items[bucketNo].next
		m.count--

		return
	}

	for item := m.items[bucketNo]; item.next != nil; item = item.next {
		if item.next.key == key {
			item.next = item.next.next
			m.count--

			return
		}
	}
}

func (m *Map[V]) Count() int {
	return m.count
}

func (m *Map[V]) rehash() {
	if len(m.items) == 0 {
		m.items = make([]*Item[V], initialSize)

		return
	}

	if float64(m.count)/float64(len(m.items)) <= maxLoadFactor {
		return
	}

	items := m.items
	m.items = make([]*Item[V], 2*len(m.items))
	for i := 0; i < len(items); i++ {
		for item := items[i]; item != nil; item = item.next {
			m.put(item.key, item.value)
		}
	}
}

func (m *Map[V]) put(key string, value V) bool {
	bucketNo := m.getBucketNo(key)

	for item := m.items[bucketNo]; item != nil; item = item.next {
		if item.key == key {
			item.value = value

			return false
		}
	}

	m.items[bucketNo] = &Item[V]{
		key:   key,
		value: value,
		next:  m.items[bucketNo],
	}

	return true
}

func (m *Map[V]) getBucketNo(key string) uint64 {
	return hash(key) % uint64(len(m.items))
}

func hash(s string) uint64 {
	sha := sha256.Sum256([]byte(s))

	return binary.LittleEndian.Uint64(sha[:])
}

package openmap

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
	key       string
	value     V
	isTrashed bool
}

func (m *Map[V]) Get(key string) V {
	v, _ := m.Find(key)

	return v
}

func (m *Map[V]) Find(key string) (V, bool) {
	var zero V
	index := m.getIndex(key)

	size := uint64(len(m.items))
	for i := uint64(0); i < size; i++ {
		offset := (index + i) % size
		if m.items[offset] == nil {
			return zero, false
		}
		if m.items[offset].isTrashed {
			continue
		}
		if m.items[offset].key == key {
			return m.items[offset].value, true
		}
	}

	return zero, false
}

func (m *Map[V]) Put(key string, value V) {
	m.rehash()

	if m.put(key, value) {
		m.count++
	}
}

func (m *Map[V]) Delete(key string) {
	index := m.getIndex(key)

	size := uint64(len(m.items))
	for i := uint64(0); i < size; i++ {
		offset := (index + i) % size
		if m.items[offset] == nil {
			break
		}
		if m.items[offset].isTrashed {
			continue
		}
		if m.items[offset].key == key {
			m.items[offset].isTrashed = true
			m.count--
			break
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
		if items[i] != nil {
			m.put(items[i].key, items[i].value)
		}
	}
}

func (m *Map[V]) put(key string, value V) bool {
	index := m.getIndex(key)

	size := uint64(len(m.items))
	for i := uint64(0); i < size; i++ {
		offset := (index + i) % size
		if m.items[offset] == nil {
			m.items[offset] = &Item[V]{key: key, value: value}
			return true
		}
		if m.items[offset].isTrashed {
			if m.items[offset].key == key {
				m.items[offset].isTrashed = false
				m.items[offset].value = value

				return true
			}
			continue
		}
		if m.items[offset].key == key {
			m.items[offset].value = value
			return false
		}
	}

	// недостижимое состояние при корректной работе карты
	panic("map items overflow")
}

func (m *Map[V]) getIndex(key string) uint64 {
	return hash(key) % uint64(len(m.items))
}

func hash(s string) uint64 {
	sha := sha256.Sum256([]byte(s))

	return binary.LittleEndian.Uint64(sha[:])
}

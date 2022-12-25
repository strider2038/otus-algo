package openmap

import (
	"crypto/sha256"
	"encoding/binary"
	"math"

	"github.com/strider2038/otus-algo/hw-03-algebraic-algos/prime"
)

const (
	maxLoadFactor = 0.5
	maxCount      = 100_000_000
)

var sizes = calcSizes()

type Map[V any] struct {
	count     int
	sizeIndex int
	items     []*Item[V]
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

	trashedOffset := uint64(0)
	trashedFound := false

	for i := uint64(0); i < uint64(len(m.items)); i++ {
		offset := m.probe(index, i)
		// если пусто, то элемента не существует
		if m.items[offset] == nil {
			return zero, false
		}

		// удаленные элементы пропускаем
		if m.items[offset].isTrashed {
			// запоминаем удаленный элемент для перестановки
			if m.items[offset].key == key {
				trashedOffset = offset
				trashedFound = true
			}
			continue
		}

		// совпадение ключа - элемент найден
		if m.items[offset].key == key {
			value := m.items[offset].value

			// оптимизация: меняем местами с удаленным
			if trashedFound {
				m.items[offset], m.items[trashedOffset] = m.items[trashedOffset], m.items[offset]
			}

			return value, true
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

	for i := uint64(0); i < uint64(len(m.items)); i++ {
		offset := m.probe(index, i)
		// если пусто, то элемента не существует
		if m.items[offset] == nil {
			return
		}

		// удаленные элементы пропускаем
		if m.items[offset].isTrashed {
			continue
		}

		if m.items[offset].key == key {
			// отмечаем элемент как удаленный
			m.items[offset].isTrashed = true
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
		m.sizeIndex = 2
		m.items = make([]*Item[V], sizes[m.sizeIndex])

		return
	}

	if float64(m.count)/float64(len(m.items)) < maxLoadFactor {
		return
	}

	items := m.items
	m.sizeIndex++
	if m.sizeIndex >= len(sizes) {
		panic("map oversize")
	}

	m.items = make([]*Item[V], sizes[m.sizeIndex])
	for i := 0; i < len(items); i++ {
		// перезаписываем в новую карту только существующие элементы
		if items[i] != nil && !items[i].isTrashed {
			m.put(items[i].key, items[i].value)
		}
	}
}

func (m *Map[V]) put(key string, value V) bool {
	index := m.getIndex(key)

	trashedOffset := uint64(0)
	trashedFound := false

	for i := uint64(0); i < uint64(len(m.items)); i++ {
		offset := m.probe(index, i)
		// если пусто, то записываем новый элемент
		if m.items[offset] == nil {
			// замена удаленного элемента
			if trashedFound {
				m.items[trashedOffset].isTrashed = false
				m.items[trashedOffset].value = value
			} else {
				m.items[offset] = &Item[V]{key: key, value: value}
			}

			return true
		}

		// если помечен как удаленный, то запоминаем его для последующей замены
		if m.items[offset].isTrashed {
			if m.items[offset].key == key {
				trashedOffset = offset
				trashedFound = true
			}
			continue
		}

		// при совпадении ключей просто обновляем значение
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

func (m *Map[V]) probe(index uint64, offset uint64) uint64 {
	return (index + offset*offset) % uint64(len(m.items))
}

func hash(s string) uint64 {
	sha := sha256.Sum256([]byte(s))

	return binary.LittleEndian.Uint64(sha[:])
}

// Размерная сетка рассчитывается на основе простых чисел. Следующее число
// из списка берется умножением индекса на два. Пример последовательности:
// 3, 7, 19, 53, 131, 311, 719, 1619, 3671, 8161, 17863, 38873, ...
func calcSizes() []int {
	primes := prime.FindBySieveOfEratosthenesOptimized(maxCount)
	n := int(math.Log2(float64(len(primes))))
	j := 0

	sizes := make([]int, n)
	for i := 2; i < len(primes); i *= 2 {
		sizes[j] = primes[i]
		j++
	}

	return sizes
}

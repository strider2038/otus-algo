package bittrie

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Array64 префиксное дерево с оптимизацией памяти и произвольным словарем символов.
// Можно индексировать до 64 символов в узле.
//
// Оптимизация заключается в том, что дочерние узлы хранятся в массиве переменной длины
// (в исходном алгоритме выделяется массив из N элементов). Для оптимизации
// доступа индексы дочерних узлов хранятся в отдельной битовой маске.
// Таким образом класс сложности доступа к дочерним узлам остается O(1) вместо O(N)
// в варианте когда необходимо перебирать все дочерние узлы.
type Array64[V any] struct {
	// Таблица индексов символов (в ключах - символ, в значении - порядковый номер)
	indices []int8
	root    array64Node[V]
	count   int
}

func NewArray64[V any](alphabet string) *Array64[V] {
	if len(alphabet) == 0 {
		panic("empty alphabet")
	}
	if len(alphabet) > 64 {
		panic("too big alphabet")
	}

	// размер индексной таблицы - наибольший символ из алфавита
	maxIndex := 0
	for _, char := range alphabet {
		if int(char) > maxIndex {
			maxIndex = int(char)
		}
	}

	indices := make([]int8, maxIndex+1)

	// заполняем таблицу специальными значениями
	for i := range indices {
		indices[i] = -1
	}
	// заполняем таблицу порядковыми номерами
	for i, char := range alphabet {
		indices[char] = int8(i)
	}

	return &Array64[V]{indices: indices}
}

func (array *Array64[V]) Count() int {
	return array.count
}

func (array *Array64[V]) Get(key string) V {
	v, _ := array.Find(key)

	return v
}

func (array *Array64[V]) Find(key string) (V, bool) {
	node := &array.root
	var zero V

	for _, char := range key {
		index := array.getCharIndex(char)
		// если индекс отсутствует в маске, то такого элемента нет в дереве
		if !node.bits.isSet(index) {
			return zero, false
		}
		// по номеру символа находим индекс следующего подузла дерева
		i := node.bits.getOneNumber(index)
		node = &node.children[i]
	}

	if node.value != nil {
		return *node.value, true
	}

	return zero, false
}

func (array *Array64[V]) Put(key string, value V) {
	node := &array.root

	for _, char := range key {
		index := array.getCharIndex(char)
		i := 0
		// если индекс найден в маске, то находим номер следующего узла в массиве
		if node.bits.isSet(index) {
			i = node.bits.getOneNumber(index)
		} else {
			// если не найден, то устанавливаем бит
			node.bits.set(index)
			// находим его порядковый номер
			i = node.bits.getOneNumber(index)
			// расширяем массив вставляя новый элемент по указанному индексу i
			node.insertChildAt(i, char)
		}
		node = &node.children[i]
	}

	// если такого элемента еще не существовало в дереве, то
	// увеличиваем счетчик количества элементов
	if node.value == nil {
		array.count++
	}

	node.value = &value
}

// Delete удаляет значение из ассоциативного массива. Реализовано в виде
// простой версии без освобождения памяти и уменьшения количества узлов.
func (array *Array64[V]) Delete(key string) {
	node := &array.root

	for _, char := range key {
		index := array.getCharIndex(char)
		// если индекс отсутствует в маске, то такого элемента нет в дереве
		if !node.bits.isSet(index) {
			return
		}
		// по номеру символа находим индекс следующего подузла дерева
		i := node.bits.getOneNumber(index)
		node = &node.children[i]
	}

	// если значение установлено для узла, то уменьшаем счетчик количества элементов
	if node.value != nil {
		array.count--
	}

	// удаляем ссылку на значение
	node.value = nil
}

// Iterate перебирает дерево и для каждого существующего узла вызывает функцию f.
func (array *Array64[V]) Iterate(f func(key string, value V) error) error {
	return array.root.iterate("", f)
}

// MemSize возвращаем примерный объем занимаемой памяти в байтах без учета хранимых значений.
func (array *Array64[V]) MemSize() int {
	return len(array.indices) +
		8 /* count */ +
		array.root.memSize()
}

func (array *Array64[V]) MarshalJSON() ([]byte, error) {
	var data bytes.Buffer
	data.WriteRune('{')
	i := 0

	err := array.Iterate(func(key string, value V) error {
		if i > 0 {
			data.WriteRune(',')
		}

		k, _ := json.Marshal(key)
		data.Write(k)

		data.WriteRune(':')
		v, err := json.Marshal(value)
		if err != nil {
			return err
		}

		data.Write(v)

		i++

		return nil
	})
	if err != nil {
		return nil, err
	}
	data.WriteRune('}')

	return data.Bytes(), nil
}

// getCharIndex возвращает порядковый номер символа из алфавитной таблицы.
func (array *Array64[V]) getCharIndex(char rune) int8 {
	if int(char) > len(array.indices) {
		panic(fmt.Sprintf("index out of range: char '%c'", char))
	}

	index := array.indices[char]
	if index < 0 {
		panic(fmt.Sprintf("index out of range: char '%c'", char))
	}

	return index
}

type array64Node[V any] struct {
	// Символ
	char rune
	// Битовая маска для индексации массива нижележащих узлов
	bits bitIndex
	// Массив нижележащих узлов переменной длины (на основе слайса)
	children []array64Node[V]
	// Ссылка на значение ассоциативного массива
	value *V
}

func (node *array64Node[V]) insertChildAt(index int, char rune) {
	n := array64Node[V]{char: char}
	if len(node.children) == index {
		// вставка в конец слайса (расширение массива)
		node.children = append(node.children, n)
		return
	}

	// вставка в середину слайса со смещением элементов > index вправо
	node.children = append(node.children[:index+1], node.children[index:]...)
	node.children[index] = n
}

func (node *array64Node[V]) iterate(key string, f func(key string, value V) error) error {
	for _, child := range node.children {
		k := key + string(child.char)
		if child.value != nil {
			if err := f(k, *child.value); err != nil {
				return err
			}
		}
		if err := child.iterate(k, f); err != nil {
			return err
		}
	}

	return nil
}

func (node *array64Node[V]) memSize() int {
	nodeSize := 4 /*char*/ +
		8 /* bits */ +
		8 /* value pointer */ +
		8 /* children pointer */
	size := nodeSize

	for _, child := range node.children {
		size += child.memSize()
	}

	size += (cap(node.children) - len(node.children)) * nodeSize

	return size
}

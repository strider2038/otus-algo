package bittrie

import "math/bits"

// bitIndex - битовая маска для хранения 64 индексов.
type bitIndex uint64

func (b *bitIndex) set(n int8) {
	*b = *b | (1 << n)
}

func (b *bitIndex) unset(n int8) {
	*b = *b & ^(1 << n)
}

func (b *bitIndex) isSet(n int8) bool {
	return *b&(1<<n) != 0
}

// getOneNumber возвращает порядковый номер установленного бита. Перед вызовом функции
// необходимо обязательно проверить установлен ли бит с помощью функции isSet.
//
// Пример маски и номеров
//
//	маска             0 0 1 0 0 1 1 0
//	номер бита        7 6 5 4 3 2 1 0
//	порядковый номер  - - 2 - - 1 0 -
//
// Примеры:
//
//	маска bitIndex = 0010 0110, номер бита n = 1, вернется число 0
//	маска bitIndex = 0010 0110, номер бита n = 2, вернется число 1
//	маска bitIndex = 0010 0110, номер бита n = 6, вернется число 2
func (b *bitIndex) getOneNumber(n int8) int {
	return bits.OnesCount64(uint64(*b) & ^(uint64(0xFFFFFFFFFFFFFFFF) << n))
}
